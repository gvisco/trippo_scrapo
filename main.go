package main

import (
	"encoding/csv"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gvisco/trippo_scrappo/pkg/tripadvisor"
)

func main() {
	// Scraper variables
	var locationURL string
	var lang string
	var outfile string

	flag.StringVar(&locationURL, "url", "", "The URL of the target Hotel")
	flag.StringVar(&lang, "lang", "en", "The language of the review like: en,  it, etc.")
	flag.StringVar(&outfile, "out", "reviews.csv", "The output CSV file")
	flag.Parse()

	if locationURL == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	// locationURL := os.Args[1]
	log.Printf("Location URL: %s", locationURL)
	log.Printf("Language: %s", lang)

	// Get the query type from the URL
	queryType := tripadvisor.GetURLType(locationURL)
	if queryType == "" {
		log.Fatal("Invalid URL")
	}
	log.Printf("Location Type: %s", queryType)

	// Parse the location ID and location name from the URL
	locationID, locationName, err := tripadvisor.ParseURL(locationURL, queryType)
	if err != nil {
		log.Fatalf("Error parsing URL: %v", err)
	}
	log.Printf("Location ID: %d", locationID)
	log.Printf("Location Name: %s", locationName)

	// Get the query ID for the given query type.
	queryID := tripadvisor.GetQueryID(queryType)
	if err != nil {
		log.Fatal("The location ID must be an positive integer")
	}

	// The default HTTP client
	log.Print("Get HTTP Client... ")
	client := &http.Client{}
	log.Println("Done (HTTP Client)")

	// Fetch the review count for the given location ID
	log.Print("Get reviews count... ")
	reviewCount, err := tripadvisor.FetchReviewCount(client, locationID, queryType, lang)
	if err != nil {
		log.Fatalf("Error fetching review count: %v", err)
	}
	if reviewCount == 0 {
		log.Fatalf("No reviews found for location ID %d", locationID)
	}
	log.Printf("Done (Review count): %d", reviewCount)

	var fileHandle *os.File
	var writer *csv.Writer
	if _, err := os.Stat(outfile); err == nil {
		log.Printf("Append to existing file: %s", outfile)
		fileHandle, err = os.OpenFile(outfile, os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("Error opening file %s: %v", outfile, err)
		}
		writer = csv.NewWriter(fileHandle)
	} else {
		log.Printf("Creating file: %s", outfile)
		fileHandle, err = os.Create(outfile)
		if err != nil {
			log.Fatalf("Error creating file %s: %v", outfile, err)
		}
		writer = csv.NewWriter(fileHandle)
		// write the headers
		headers := []string{"Location Name", "Lang", "Title", "Text", "Rating", "Year", "Month", "Day"}
		err = writer.Write(headers)
		if err != nil {
			log.Fatalf("Error writing header to csv: %v", err)
		}
	}
	defer fileHandle.Close()

	// Calculate the number of iterations required to fetch all reviews
	iterations := tripadvisor.CalculateIterations(uint32(reviewCount))
	log.Printf("Total Iterations: %d", iterations)

	// Create a slice to store the data to be written to the CSV file
	dataToWrite := make([][]string, 0, reviewCount)

	// Scrape the reviews
	for i := uint32(0); i < iterations; i++ {

		// Introduce random delay to avoid getting blocked. The delay is between 1 and 5 seconds
		delay := rand.Intn(5) + 1
		log.Printf("Iteration: %d. Delaying for %d seconds", i, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		// Calculate the offset for the current iteration
		offset := tripadvisor.CalculateOffset(i)

		// Make the request to the TripAdvisor GraphQL endpoint
		resp, err := tripadvisor.MakeRequest(client, queryID, lang, locationID, offset, 20)
		if err != nil {
			log.Fatalf("Error making request at iteration %d: %v", i, err)
		}

		// Check if responses is nil before dereferencing
		if resp == nil {
			log.Fatalf("Received nil response for location ID %d at iteration: %d", locationID, i)
		}

		// Now it's safe to dereference responses
		response := *resp

		// Check if the response is not empty and if the response contains reviews
		if len(response) > 0 && len(response[0].Data.Locations) > 0 {

			// Get the reviews from the response
			reviews := response[0].Data.Locations[0].ReviewListPage.Reviews

			// Iterating over the reviews
			for _, row := range reviews {
				row := []string{
					locationName,
					lang,
					row.Title,
					row.Text,
					strconv.Itoa(row.Rating),
					row.CreatedDate[0:4],
					row.CreatedDate[5:7],
					row.CreatedDate[8:10],
				}

				// Append the row to the dataToWrite slice
				dataToWrite = append(dataToWrite, row)
			}

		}

	}

	// Write data to the CSV file
	err = writer.WriteAll(dataToWrite)
	if err != nil {
		log.Fatalf("Error writing data to csv: %v", err)
	}
}

# trippo_scrappo
A simple command-line scraper for TripAdvisor reviews, written in GO.

## Info
Heavily based on the code from [TripAdvisor-Review-Scraper](https://github.com/algo7/TripAdvisor-Review-Scraper) (that works like a charm by the way), I have modified it to get to a simple command-line experience, getting rid of all the docker-related features and and adding a bit of flexibility where needed.

Main changes include
- Input arguments via command line instead of environment variables
- Support to scrape reviews in different laguages
- Append the output to an existing CSV file, if the file already exists

## How to use it
It requires Go +v1.21.

Run the `main` to see the help
```bash
$ go run main.go 
  -lang string
        The language of the review like: en,  it, etc. (default "en")
  -out string
        The output CSV file (default "reviews.csv")
  -url string
        The URL of the target Hotel
```
The `url` is mandatory and is the URL of the TripAdvisor page of the hotel/restaurant/airline you want to scrape from `https://www.tripadvisor.com` (note that other domains, different from `.com` are not supported). 
The URL should be in the following format:
1. Airline: `https://www.tripadvisor.com/Airline_Review-d8729113-Reviews-Lufthansa`
2. Hotel: `https://www.tripadvisor.com/Hotel_Review-g188107-d231860-Reviews-Beau_Rivage_Palace-Lausanne_Canton_of_Vaud.html`
3. Restaurant: `https://www.tripadvisor.com/Restaurant_Review-g187265-d11827759-Reviews-La_Terrasse-Lyon_Rhone_Auvergne_Rhone_Alpes.html`

### Example
This example shows how to download all the reviews, in French, of the [Beau-Rivage Palace](https://www.tripadvisor.com/Hotel_Review-g188107-d231860-Reviews-Beau_Rivage_Palace-Lausanne_Canton_of_Vaud.html) hotel. The output will be saved in the `BeauRivage.CSV` file.

```bash
$go run main.go -url=https://www.tripadvisor.com/Hotel_Review-g188107-d231860-Reviews-Beau_Rivage_Palace-Lausanne_Canton_of_Vaud.html -lang=fr -out="BeauRivage.csv"
2024/02/04 17:43:34 Location URL: https://www.tripadvisor.com/Hotel_Review-g188107-d231860-Reviews-Beau_Rivage_Palace-Lausanne_Canton_of_Vaud.html
2024/02/04 17:43:34 Language: fr
2024/02/04 17:43:34 Location Type: HOTEL
2024/02/04 17:43:34 Location ID: 231860
2024/02/04 17:43:34 Location Name: Beau_Rivage_Palace
2024/02/04 17:43:34 Get HTTP Client...
2024/02/04 17:43:34 Done (HTTP Client)
2024/02/04 17:43:34 Get reviews count...
2024/02/04 17:43:34 Sending request...
2024/02/04 17:43:35 Done (Response received)
2024/02/04 17:43:35 Done (Review count): 694
2024/02/04 17:43:35 Creating file: BeauRivage
2024/02/04 17:43:35 Total Iterations: 35
2024/02/04 17:43:35 Iteration: 0. Delaying for 4 seconds
2024/02/04 17:43:39 Sending request... 
2024/02/04 17:43:39 Done (Response received)
2024/02/04 17:43:39 Iteration: 1. Delaying for 5 seconds
2024/02/04 17:43:44 Sending request... 
2024/02/04 17:43:44 Done (Response received)
2024/02/04 17:43:44 Iteration: 2. Delaying for 3 seconds
2024/02/04 17:43:47 Sending request... 
2024/02/04 17:43:48 Done (Response received)
[...]
2024/02/04 17:45:38 Iteration: 34. Delaying for 3 seconds
2024/02/04 17:45:41 Sending request... 
2024/02/04 17:45:41 Done (Response received)
```

## License and Credits
This software is based on the [TripAdvisor-Review-Scraper](https://github.com/algo7/TripAdvisor-Review-Scraper) and it is released accordingly under a GNU license. If you want to credit the authors, please refer to the original [CITATION](https://github.com/algo7/TripAdvisor-Review-Scraper/blob/main/CITATION.cff) instructions.
# trippo_scrappo
A simple command-line scraper for TripAdvisor reviews, written in GO.

## Info
Heavily based on the code from [TripAdvisor-Review-Scraper](https://github.com/gvisco/TripAdvisor-Review-Scraper) (that works like a charm by the way), I have modified it to get to a simple command-line experience, getting rid of all the docker-related features and and adding a bit of flexibility where needed.

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
2024/02/04 17:43:48 Iteration: 3. Delaying for 5 seconds
2024/02/04 17:43:53 Sending request... 
2024/02/04 17:43:53 Done (Response received)
2024/02/04 17:43:53 Iteration: 4. Delaying for 2 seconds
2024/02/04 17:43:55 Sending request... 
2024/02/04 17:43:56 Done (Response received)
2024/02/04 17:43:56 Iteration: 5. Delaying for 2 seconds
2024/02/04 17:43:58 Sending request... 
2024/02/04 17:43:59 Done (Response received)
2024/02/04 17:43:59 Iteration: 6. Delaying for 3 seconds
2024/02/04 17:44:02 Sending request... 
2024/02/04 17:44:02 Done (Response received)
2024/02/04 17:44:02 Iteration: 7. Delaying for 3 seconds
2024/02/04 17:44:05 Sending request... 
2024/02/04 17:44:05 Done (Response received)
2024/02/04 17:44:05 Iteration: 8. Delaying for 3 seconds
2024/02/04 17:44:08 Sending request... 
2024/02/04 17:44:09 Done (Response received)
2024/02/04 17:44:09 Iteration: 9. Delaying for 4 seconds
2024/02/04 17:44:13 Sending request... 
2024/02/04 17:44:13 Done (Response received)
2024/02/04 17:44:13 Iteration: 10. Delaying for 4 seconds
2024/02/04 17:44:17 Sending request... 
2024/02/04 17:44:17 Done (Response received)
2024/02/04 17:44:17 Iteration: 11. Delaying for 5 seconds
2024/02/04 17:44:22 Sending request... 
2024/02/04 17:44:23 Done (Response received)
2024/02/04 17:44:23 Iteration: 12. Delaying for 5 seconds
2024/02/04 17:44:28 Sending request... 
2024/02/04 17:44:28 Done (Response received)
2024/02/04 17:44:28 Iteration: 13. Delaying for 4 seconds
2024/02/04 17:44:32 Sending request... 
2024/02/04 17:44:33 Done (Response received)
2024/02/04 17:44:33 Iteration: 14. Delaying for 3 seconds
2024/02/04 17:44:36 Sending request... 
2024/02/04 17:44:36 Done (Response received)
2024/02/04 17:44:36 Iteration: 15. Delaying for 1 seconds
2024/02/04 17:44:37 Sending request... 
2024/02/04 17:44:38 Done (Response received)
2024/02/04 17:44:38 Iteration: 16. Delaying for 2 seconds
2024/02/04 17:44:40 Sending request... 
2024/02/04 17:44:40 Done (Response received)
2024/02/04 17:44:40 Iteration: 17. Delaying for 1 seconds
2024/02/04 17:44:41 Sending request... 
2024/02/04 17:44:41 Done (Response received)
2024/02/04 17:44:41 Iteration: 18. Delaying for 4 seconds
2024/02/04 17:44:45 Sending request... 
2024/02/04 17:44:46 Done (Response received)
2024/02/04 17:44:46 Iteration: 19. Delaying for 4 seconds
2024/02/04 17:44:50 Sending request... 
2024/02/04 17:44:50 Done (Response received)
2024/02/04 17:44:50 Iteration: 20. Delaying for 5 seconds
2024/02/04 17:44:55 Sending request... 
2024/02/04 17:44:55 Done (Response received)
2024/02/04 17:44:55 Iteration: 21. Delaying for 2 seconds
2024/02/04 17:44:57 Sending request... 
2024/02/04 17:44:58 Done (Response received)
2024/02/04 17:44:58 Iteration: 22. Delaying for 2 seconds
2024/02/04 17:45:00 Sending request... 
2024/02/04 17:45:00 Done (Response received)
2024/02/04 17:45:00 Iteration: 23. Delaying for 4 seconds
2024/02/04 17:45:04 Sending request... 
2024/02/04 17:45:05 Done (Response received)
2024/02/04 17:45:05 Iteration: 24. Delaying for 5 seconds
2024/02/04 17:45:10 Sending request... 
2024/02/04 17:45:10 Done (Response received)
2024/02/04 17:45:10 Iteration: 25. Delaying for 5 seconds
2024/02/04 17:45:15 Sending request... 
2024/02/04 17:45:16 Done (Response received)
2024/02/04 17:45:16 Iteration: 26. Delaying for 1 seconds
2024/02/04 17:45:17 Sending request... 
2024/02/04 17:45:17 Done (Response received)
2024/02/04 17:45:17 Iteration: 27. Delaying for 4 seconds
2024/02/04 17:45:21 Sending request... 
2024/02/04 17:45:21 Done (Response received)
2024/02/04 17:45:21 Iteration: 28. Delaying for 1 seconds
2024/02/04 17:45:22 Sending request... 
2024/02/04 17:45:23 Done (Response received)
2024/02/04 17:45:23 Iteration: 29. Delaying for 2 seconds
2024/02/04 17:45:25 Sending request... 
2024/02/04 17:45:25 Done (Response received)
2024/02/04 17:45:25 Iteration: 30. Delaying for 1 seconds
2024/02/04 17:45:26 Sending request... 
2024/02/04 17:45:27 Done (Response received)
2024/02/04 17:45:27 Iteration: 31. Delaying for 3 seconds
2024/02/04 17:45:30 Sending request... 
2024/02/04 17:45:30 Done (Response received)
2024/02/04 17:45:30 Iteration: 32. Delaying for 2 seconds
2024/02/04 17:45:32 Sending request... 
2024/02/04 17:45:33 Done (Response received)
2024/02/04 17:45:33 Iteration: 33. Delaying for 5 seconds
2024/02/04 17:45:38 Sending request... 
2024/02/04 17:45:38 Done (Response received)
2024/02/04 17:45:38 Iteration: 34. Delaying for 3 seconds
2024/02/04 17:45:41 Sending request... 
2024/02/04 17:45:41 Done (Response received)
```
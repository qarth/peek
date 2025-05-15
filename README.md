# Seek Job Scraper

This is a Go application that scrapes job listings from seek.com.au using goquery and prints them to the console.

## Features
- Scrapes job title, company, location, and link from Seek job listings
- Prints results in a readable format

## Requirements
- Go 1.18 or newer
- Internet connection

## Usage
1. Clone this repository or copy the code into your Go workspace.
2. Install dependencies:
   ```bash
   go get github.com/PuerkitoBio/goquery
   ```
3. Run the application:
   ```bash
   go run main.go
   ```

You can modify the `jobTitle` and `location` variables in `main.go` to change the search query.

## License
MIT

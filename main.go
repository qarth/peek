package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func scrapeSeek(jobTitle, location string) {
	url := fmt.Sprintf("https://www.seek.com.au/%s-jobs/in-%s", jobTitle, location)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error performing request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: Received status code %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Error parsing HTML: %v", err)
	}

	doc.Find("article[data-automation='normalJob']").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a[data-automation='jobTitle']").Text()
		link, _ := s.Find("a[data-automation='jobTitle']").Attr("href")
		company := s.Find("a[data-automation='jobCompany']").Text()
		location := s.Find("a[data-automation='jobLocation']").Text()
		seenLast := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("Job %d:\n", i+1)
		fmt.Printf("  Title: %s\n", strings.TrimSpace(title))
		fmt.Printf("  Link: https://www.seek.com.au%s\n", link)
		fmt.Printf("  Company: %s\n", strings.TrimSpace(company))
		fmt.Printf("  Location: %s\n", strings.TrimSpace(location))
		fmt.Printf("  Seen: %s\n", seenLast)
		fmt.Println()
	})
}

func main() {
	jobTitle := "mining engineer"
	location := "All Australia"
	scrapeSeek(jobTitle, location)
}

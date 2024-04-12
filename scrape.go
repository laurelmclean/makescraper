package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Concert struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// Slice to store concerts
	var concerts []Concert

	c.OnHTML("div.elementor-loop-container.elementor-grid", func(e *colly.HTMLElement) {
		name := e.ChildText("div.elementor-widget-theme-post-title a")
		date := e.ChildText("div.elementor-widget-shortcode > div > div > div")
		// create a new concert struct
		concert := Concert{name, date}
		// append the concert struct to the slice
		concerts = append(concerts, concert)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// After scraping is complete
	c.OnScraped(func(r *colly.Response) {
		// turn the slice into a JSON string
		concertsJson, err := json.Marshal(concerts)
		if err != nil {
			log.Fatalf("Failed to convert to JSON: %v", err)
		}
		// create file
		file, err := os.Create("concerts.json")
		if err != nil {
			log.Fatalf("Failed to create file: %v", err)
		}
		defer file.Close()

		// add json to the file
		writer := bufio.NewWriter(file)
		_, err = writer.WriteString(string(concertsJson))
		if err != nil {
			log.Fatalf("Failed to write json to file: %v", err)
		}
		// flush the buffer to ensure data is written to the file
		writer.Flush()

		fmt.Println("Data written to concerts.json successfully.")
	})

	// Start scraping on https://thepalomino.ca/live-events/
	c.Visit("https://thepalomino.ca/live-events/")
}

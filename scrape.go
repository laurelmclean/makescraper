package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

// Concert struct for each concert
type Concert struct {
	Name string
	Date string
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// Slice to store concerts
	var concerts []Concert

	// // On every div element which has the specified class for concert name, call callback for concert name
	// c.OnHTML("div.elementor-loop-container.elementor-grid div.elementor-widget-theme-post-title a", func(e *colly.HTMLElement) {
	// 	concert := Concert{
	// 		Name: e.Text,
	// 	}
	// 	concerts = append(concerts, concert)
	// })

	// // On every div element which has the specified class for concert date, call callback for concert date
	// c.OnHTML("div.elementor-loop-container.elementor-grid div.elementor-widget-shortcode > div > div > div", func(e *colly.HTMLElement) {
	// 	concert := Concert{
	// 		Date: e.Text,
	// 	}
	// 	concerts = append(concerts, concert)
	// })

	c.OnHTML("div.elementor-loop-container.elementor-grid", func(e *colly.HTMLElement) {
		name := e.ChildText("div.elementor-widget-theme-post-title a")
		date := e.ChildText("div.elementor-widget-shortcode > div > div > div")
		concert := Concert{
			Name: name,
			Date: date,
		}
		concerts = append(concerts, concert)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://thepalomino.ca/live-events/
	c.Visit("https://thepalomino.ca/live-events/")

	// Print the scraped data
	for _, concert := range concerts {
		fmt.Printf("Concert Name: %s\n", concert.Name)
		fmt.Printf("Concert Date: %s\n", concert.Date)
		fmt.Println("------------")
	}
}

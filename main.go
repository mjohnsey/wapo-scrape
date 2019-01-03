package main

import (
	"encoding/json"
	"os"

	"github.com/gocolly/colly"
)

// Headline represents a news headline
type Headline struct {
	URL   *string `json:"url"`
	Title *string `json:"title"`
	Blurb *string `json:"blurb,omitempty"`
}

type WashingtonPostScraper struct {
	collector *colly.Collector
}

func (s WashingtonPostScraper) CreateNewScraper() *WashingtonPostScraper {
	c := colly.NewCollector()
	// c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))
	c.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"

	scraper := WashingtonPostScraper{
		collector: c,
	}
	return &scraper
}

func (s WashingtonPostScraper) ScrapeHeadlines() (*[]*Headline, error) {
	headlines := make([]*Headline, 0)
	s.collector.OnHTML("#main-content", func(e *colly.HTMLElement) {
		e.ForEach("div.flex-item", func(ndx int, flex *colly.HTMLElement) {
			url := flex.ChildAttr("div.headline > a", "href")
			if url != "" {
				title := flex.ChildText("div.headline > a")
				blurb := flex.ChildText("div.blurb")
				newHeadline := &Headline{
					URL:   &url,
					Title: &title,
				}
				if blurb != "" {
					newHeadline.Blurb = &blurb
				}

				headlines = append(headlines, newHeadline)
			}
		})
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL.String())
	// })

	s.collector.Visit("https://www.washingtonpost.com/?noredirect=on")
	return &headlines, nil
}

func scrapeWashingtonPostHeadlines() (*[]*Headline, error) {
	scraper := WashingtonPostScraper{}.CreateNewScraper()
	return scraper.ScrapeHeadlines()
}

func main() {

	headlines, _ := scrapeWashingtonPostHeadlines()
	enc := json.NewEncoder(os.Stdout)

	// Dump json to the standard output
	enc.Encode(headlines)
}

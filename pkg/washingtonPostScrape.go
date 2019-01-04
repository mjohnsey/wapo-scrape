package pkg

import (
	"strconv"
	"time"
)

// WashingtonPostScrape represents the result of scraping the Washington Post for headlines
type WashingtonPostScrape struct {
	Headlines  *[]*Headline `json:"headlines,omitempty"`
	ScrapeTime string       `json:"scrapeTime"`
}

// SetTimeToNow sets the scrape's timestamp to now in unix time
func (s *WashingtonPostScrape) SetTimeToNow() {
	s.ScrapeTime = strconv.FormatInt(time.Now().Unix(), 10)
}

// ScrapeWashingtonPost is a runner for scraping the Washington Post
func (s WashingtonPostScrape) ScrapeWashingtonPost() (*WashingtonPostScrape, error) {
	scraper := WashingtonPostScraper{}.CreateNewWashingtonPostScraper()
	headlines, err := scraper.ScrapeHeadlines()
	if err != nil {
		return nil, err
	}
	scrape := WashingtonPostScrape{
		Headlines: headlines,
	}
	scrape.SetTimeToNow()
	return &scrape, nil
}

package pkg

import (
	"time"

	"github.com/gocolly/colly"
)

// WashingtonPostScraper is a wrapper around the scraper
type WashingtonPostScraper struct {
	collector *colly.Collector
}

// URL returns the WaPo homepage
func (s WashingtonPostScraper) URL() string {
	return "https://www.washingtonpost.com/?noredirect=on"
}

// UserAgent returns the user agent used to scrape the Washington Post
func (s WashingtonPostScraper) UserAgent() string {
	// WaPo does some weird filtering if it doesn't think this is a browser, so I feeling lucky'd this one
	// TODO: replace with a better UA
	return "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"
}

// CreateNewWashingtonPostScraper is a constructor for a WashingtonPostScraper
func (s WashingtonPostScraper) CreateNewWashingtonPostScraper() *WashingtonPostScraper {
	c := colly.NewCollector()
	// c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))
	c.UserAgent = s.UserAgent()
	c.IgnoreRobotsTxt = false

	// Adding this wait so AJAX can load, might need to look at https://github.com/chromedp/chromedp in the future
	c.Limit(&colly.LimitRule{
		Delay: 5 * time.Second,
	})

	scraper := WashingtonPostScraper{
		collector: c,
	}
	return &scraper
}

// ScrapeHeadlines does the heavy lifting of grabbing the headlines from the Washington Post
func (s WashingtonPostScraper) ScrapeHeadlines() (*[]*Headline, error) {
	headlines := make([]*Headline, 0)
	var parseErr error
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

	// Set error handler
	s.collector.OnError(func(r *colly.Response, err error) {
		parseErr = err
	})

	s.collector.Visit(s.URL())
	return &headlines, parseErr
}

// Copyright © 2019 Michael Johnsey <mjohnsey@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
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
func ScrapeWashingtonPost() (*WashingtonPostScrape, error) {
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

// Headline represents a news headline
type Headline struct {
	URL   *string `json:"url"`
	Title *string `json:"title"`
	Blurb *string `json:"blurb,omitempty"`
}

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

// scrapeCmd represents the scrape command
var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape the Washington Post headlines",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := ScrapeWashingtonPost()
		if err != nil {
			log.Fatalln(err)
		}

		// Dump json to the standard output
		json.NewEncoder(os.Stdout).Encode(s)
	},
}

func init() {
	rootCmd.AddCommand(scrapeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scrapeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scrapeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

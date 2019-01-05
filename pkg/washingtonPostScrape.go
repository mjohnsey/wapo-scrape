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
package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

// WashingtonPostScrape represents the result of scraping the Washington Post for headlines
type WashingtonPostScrape struct {
	Headlines  *[]*Headline `json:"headlines,omitempty"`
	ScrapeTime UTCTimestamp `json:"scrapeTime"`
}

// SetTimeToNow sets the scrape's timestamp to now in unix time
func (s *WashingtonPostScrape) SetTimeToNow() {
	s.ScrapeTime = UTCTimestamp{time.Now().UTC()}
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

func (s WashingtonPostScrape) FromJsonFile(fileLocation string) (*WashingtonPostScrape, error) {
	file, fileReadErr := ioutil.ReadFile(fileLocation)
	if fileReadErr != nil {
		return nil, fileReadErr
	}
	var scrape WashingtonPostScrape
	jsonParseErr := json.Unmarshal(file, &scrape)
	if jsonParseErr != nil {
		return nil, jsonParseErr
	}
	return &scrape, nil
}

func (scrape *WashingtonPostScrape) Stats() string {
	numOfScrapes := len(*scrape.Headlines)
	stats := fmt.Sprintf("Num of headlines: %d", numOfScrapes)
	stats = fmt.Sprintf("%s\nTime: %s", stats, scrape.ScrapeTime)
	if numOfScrapes > 0 {
		top5 := 5
		if numOfScrapes < 5 {
			top5 = numOfScrapes
		}
		stats = fmt.Sprintf("%s\nTop %d Headlines:\n", stats, top5)
		headlines := *scrape.Headlines
		for i := 0; i < top5; i++ {
			stats = fmt.Sprintf("%s\n- %s (%s)", stats, *headlines[i].Title, *headlines[i].URL)
		}
	}
	return stats
}

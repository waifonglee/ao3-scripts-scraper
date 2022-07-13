package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
	"strings"
	"time"
)

func createCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("archiveofourown.org"),
	)
	c.Limit(&colly.LimitRule{
		// Filter domains affected by this rule
		DomainGlob:  "archiveofourown.org/*",
		// Set a delay between requests to these domains
		Delay: REQUEST_DELAY,
	})
	
	return c
}

func fetchSingleDownloadDetails(c *colly.Collector, ao3Url string, downloadDetails *DownloadDetails) {

	c.OnRequest(func(req *colly.Request) {
		fmt.Printf(INFO_VISITING, req.URL.String())
	})

	c.OnError(func(resp *colly.Response, err error) {
		if resp.StatusCode == http.StatusTooManyRequests {
			fmt.Println(INFO_TOO_MANY_REQ)
			time.Sleep(RETRY_TIMEOUT)
			resp.Request.Retry()
		} else {
			fetchError := &FetchError{url: resp.Request.URL.String(), statusCode: resp.StatusCode, err: err}
			panic(fetchError)
		}
	})

	c.OnHTML("li.download", func(e *colly.HTMLElement) {
		urls := make(map[string]string)

		e.ForEach("a[href]", func(_ int, child *colly.HTMLElement) {
			if child.Attr("href") == "#" {
				return
			}
			fileType := strings.TrimSpace(child.Text)
			urls[strings.ToLower(fileType)] = child.Attr("href")
		})

		downloadDetails.urls = urls
	})

	c.OnHTML("h2.title.heading", func(e *colly.HTMLElement) {
		downloadDetails.title = strings.TrimSpace(e.Text)
	})

	c.Visit(ao3Url)
}

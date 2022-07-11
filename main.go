package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type Fanfic struct {
	id int
	kudos int
	author string
	title string
	fandom string
	relationship []string
}


func main() {
	url := parseArgs()

	c := colly.NewCollector(
		colly.AllowedDomains("archiveofourown.org"),
	)

	c.Limit(&colly.LimitRule{
		// Filter domains affected by this rule
		DomainGlob:  "archiveofourown.org/*",
		// Set a delay between requests to these domains
		Delay: 5 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnHTML("li.download", func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a[href]", "href")
		fmt.Println(links)
		//downloadFile("pdf", fmt.Sprintf("https://archiveofourown.org/downloads/40144401/Klandestin.pdf?updated_at=1657212088"))
	})

	/*
	c.OnHTML("h2.title.heading", func(e *colly.HTMLElement) {
		title := e.Text
		fmt.Println(title)
		fmt.Println("hello2")
	})

	c.OnHTML("div.work", func(e *colly.HTMLElement) {
		fmt.Println("hello3")
	})
*/
	c.Visit(url)

}

func downloadFile(fileType string, fileUrl string) {
	fmt.Println(fileUrl)
	outFile, err := os.Create(fmt.Sprintf("test.%s", fileType))
	if err != nil {
		fmt.Println("File creation:", fileUrl, "failed with error:", err)
		return
	}
	defer outFile.Close()
	resp, err := http.Get(fileUrl)
	if err != nil {
		fmt.Println("File download:", fileUrl, "failed with error:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("File download:", fileUrl, "failed with status code:", resp.StatusCode)
	}
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		fmt.Println("File write:", fileUrl, "failed with error:", err)
	}
}

func parseArgs() string {
	var url string
	flag.StringVar(&url, "url", "", "URL of fanfic to download")
	flag.Parse()
	if len(url) == 0 {
		flag.PrintDefaults()
	}

	fmt.Printf("Url: %s \n", url)
	
	return url
}

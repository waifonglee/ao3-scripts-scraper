package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	ao3Url, format := parseDownloaderArgs()
	fmt.Println(ao3Url, format)

	collector := createCollector()
	downloadDetails := new(DownloadDetails)
	
	fetchSingleDownloadDetails(collector, ao3Url, downloadDetails)
	fmt.Println("download details", downloadDetails)
	time.Sleep(REQUEST_DELAY)
	title := replaceSpaceWithUnderscore(downloadDetails.title)
	fileName := title + "." + format
	dir := filepath.Join(DOWNLOAD_DIR, fileName)
	downloadUrl :=  "https://archiveofourown.org" + downloadDetails.getUrlByFormat(format)
	err := downloadFic(downloadUrl, dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


/*
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
		if r.StatusCode == http.StatusTooManyRequests {
			time.Sleep(300 * time.Second)
			r.Request.Retry()
		} else {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		}
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

*/

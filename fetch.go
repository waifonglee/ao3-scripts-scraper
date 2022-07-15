package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

var worksCollector *colly.Collector
var workCollector *colly.Collector 
var pageDetails *PageDetails
var downloadDetails *DownloadDetails


func setDownloadDetails(format string) {
	downloadDetails = &DownloadDetails{format: format}

}

func setPageDetails(endPage int) {
	pageDetails = &PageDetails{end: endPage}
}

func downloadWorkFromLink(link string, format string) {
	setDownloadDetails(format)
	setCollectorForSingleDownload()
	workCollector.Visit(link)
}

func downloadWorksFromLink(link string, endPage int, format string) {
	setPageDetails(endPage)
	setDownloadDetails(format)
	setCollectorForPages()
	setCollectorForSingleDownload()
	worksCollector.Visit(link)
}

func createCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(DOMAIN),
		colly.CacheDir(CACHE_DIR),

	)
	c.Limit(&colly.LimitRule{
		// Filter domains affected by this rule
		DomainGlob: "*",
		// Set a delay between requests to these domains
		Delay: REQUEST_DELAY,
	})

	extensions.RandomUserAgent(c)
	
	return c
}

func onReq(req *colly.Request) {
	fmt.Printf(INFO_VISITING, getFormattedTime(), req.URL.String())
}

func onResp(resp *colly.Response) {
	fmt.Println("Response code", resp.StatusCode)
}

func onErr(resp *colly.Response, err error) {
	if resp.StatusCode == http.StatusTooManyRequests {
		fmt.Printf(INFO_TOO_MANY_REQ, getFormattedTime())
		time.Sleep(RETRY_TIMEOUT)
		resp.Request.Retry()
	} else {
		fetchError := &FetchError{url: resp.Request.URL.String(), statusCode: resp.StatusCode, err: err}
		panic(fetchError) //Not sure if its a good idea to panic here but IDK man
	}
}

func setDownloadLinkFromWork(e *colly.HTMLElement) {
	e.ForEach("a[href]", func(_ int, child *colly.HTMLElement) {
		if downloadDetails.format != strings.ToLower(strings.TrimSpace(child.Text)) {
			return
		}
		downloadDetails.downloadLink = e.Request.AbsoluteURL(child.Attr("href"))
	})
}

func setTitleFromWork(e *colly.HTMLElement) {
	downloadDetails.title = strings.TrimSpace(e.Text)
}

func downloadWorkAfterFetch(resp *colly.Response) {
	if len(downloadDetails.downloadLink) == 0 || len(downloadDetails.title) == 0 {
		fmt.Printf(INFO_UNABLE_TO_DOWNLOAD, resp.Request.URL)
		return
	}
	fmt.Printf(INFO_DOWNLOADING, getFormattedTime(), resp.Request.URL)
	downloadErr := downloadSingleWork(downloadDetails)
	downloadDetails.resetTitleLink()
	defer time.Sleep(REQUEST_DELAY)
	if downloadErr != nil {
		fmt.Println(downloadErr)
		return
	}
}

func setCollectorForSingleDownload() {
	if worksCollector != nil {
		workCollector = worksCollector.Clone()
	} else {
		workCollector = createCollector()
	} 

	workCollector.OnRequest(onReq)
	workCollector.OnResponse(onResp)
	workCollector.OnError(onErr)
	workCollector.OnHTML("li.download", setDownloadLinkFromWork)
	workCollector.OnHTML("h2.title.heading", setTitleFromWork)
	workCollector.OnScraped(downloadWorkAfterFetch)
}

func goNextPage(e *colly.HTMLElement) {
	//Pagination on top and bottom in ao3, only collect from one
	if e.Index > 0 {
		return
	}

	currentPage, _ := strconv.Atoi(e.ChildText("span.current"))
	if currentPage == pageDetails.end {
		return
	}

	nextPageLink :=  e.Request.AbsoluteURL(e.ChildAttr("[rel='next']", "href"))
	//check if its the final page
	if len(nextPageLink) == 0 {
		return
	}
	e.Request.Visit(nextPageLink)
}

// Not quite sure about this function name, can't think of anything better yet
func fetchWork(e *colly.HTMLElement) { 
	link := e.Request.AbsoluteURL(e.ChildAttr("h4.heading > a", "href")) + "?view_adult=true"
	if len(link) == 0 {
		return
	}
	workCollector.Visit(link)
}

func setCollectorForPages() {
	worksCollector = createCollector()

	worksCollector.OnRequest(onReq)
	worksCollector.OnResponse(onResp)
	worksCollector.OnError(onErr)
	worksCollector.OnHTML("ol.pagination.actions", goNextPage)
	worksCollector.OnHTML("li[role='article']", fetchWork)
}
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

var worksCollector *colly.Collector
var workCollector *colly.Collector 
var pageDetails *PageDetails
var downloadDetails *DownloadDetails

func createCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(DOMAIN),
		colly.CacheDir(CACHE_DIR),

	)
	c.Limit(&colly.LimitRule{
		// Filter domains affected by this rule
		DomainGlob: fmt.Sprintf("%s/*", DOMAIN),
		// Set a delay between requests to these domains
		Delay: REQUEST_DELAY,
	})
	
	return c
}

func onReq(req *colly.Request) {
	fmt.Printf(INFO_VISITING, time.Now().String(), req.URL.String())
}

func onErr(resp *colly.Response, err error) {
	if resp.StatusCode == http.StatusTooManyRequests {
		fmt.Printf(INFO_TOO_MANY_REQ, time.Now().String())
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
		downloadDetails.link = e.Request.AbsoluteURL(child.Attr("href"))
	})
}

func setTitleFromWork(e *colly.HTMLElement) {
	downloadDetails.format = strings.TrimSpace(e.Text)
}

func setDownloadCollector() {
	if worksCollector != nil {
		workCollector = worksCollector.Clone()
	} else {
		workCollector = createCollector()
	} 

	workCollector.OnRequest(onReq)
	workCollector.OnError(onErr)
	workCollector.OnHTML("li.download", setDownloadLinkFromWork)
	workCollector.OnHTML("h2.title.heading", setTitleFromWork)
}

func setPageCollector() {
	worksCollector = createCollector()
	workCollector.OnRequest(onReq)
	workCollector.OnError(onErr)
}
func setmainCollector(endPage int) {
	mainCollector = createCollector()
	mainCollector.OnRequest(func(req *colly.Request) {
		fmt.Printf(INFO_VISITING, req.URL.String())
	})

	mainCollector.OnError(func(resp *colly.Response, err error) {
		if resp.StatusCode == http.StatusTooManyRequests {
			fmt.Println(INFO_TOO_MANY_REQ)
			time.Sleep(RETRY_TIMEOUT)
			resp.Request.Retry()
		} else {
			fetchError := &FetchError{url: resp.Request.URL.String(), statusCode: resp.StatusCode, err: err}
			panic(fetchError) //Not sure if its a good idea to panic here but IDK man
		}
	})

	mainCollector.OnHTML("ol.pagination.actions", func (e *colly.HTMLElement) {
		if e.Index > 0 {
			return
		}

		currentPage, err := strconv.Atoi(e.ChildText("span.current"))
		if err != nil {
			panic(err)
		}
		fmt.Println("current page", currentPage)

		if currentPage == endPage {
			return
		}

		nextPageUrl :=  e.Request.AbsoluteURL(e.ChildAttr("[rel='next']", "href"))
		if len(nextPageUrl) == 0 {
			return
		}
		time.Sleep(REQUEST_DELAY)
		e.Request.Visit(nextPageUrl)
	})

	mainCollector.OnHTML("li[role='article']", func (e *colly.HTMLElement) {
			link:=e.Request.AbsoluteURL(e.ChildAttr("a[href^='/works/']", "href"))
			fmt.Println(link)

			//downloadDetails := new(DownloadDetails)
			
			
			//fetchSingleDownloadDetailss(dlCollector)
			dlCollector.Visit(link)
			//fmt.Println(downloadDetails)

			/*
			filePath := formatPath(downloadDetails.title, "pdf")
			downloadUrl := fmt.Sprintf("https://%s%s", DOMAIN, downloadDetails.getUrlByFormat("pdf"))
			downloadErr := downloadSingleFic(downloadUrl, filePath)
			if downloadErr != nil {
				fmt.Println(downloadErr)
				return
			}
			
	})
}

/*
func setmainCollector(endPage int) {
	mainCollector = createCollector()
	mainCollector.OnRequest(func(req *colly.Request) {
		fmt.Printf(INFO_VISITING, req.URL.String())
	})

	mainCollector.OnError(func(resp *colly.Response, err error) {
		if resp.StatusCode == http.StatusTooManyRequests {
			fmt.Println(INFO_TOO_MANY_REQ)
			time.Sleep(RETRY_TIMEOUT)
			resp.Request.Retry()
		} else {
			fetchError := &FetchError{url: resp.Request.URL.String(), statusCode: resp.StatusCode, err: err}
			panic(fetchError) //Not sure if its a good idea to panic here but IDK man
		}
	})

	mainCollector.OnHTML("ol.pagination.actions", func (e *colly.HTMLElement) {
		if e.Index > 0 {
			return
		}

		currentPage, err := strconv.Atoi(e.ChildText("span.current"))
		if err != nil {
			panic(err)
		}
		fmt.Println("current page", currentPage)

		if currentPage == endPage {
			return
		}

		nextPageUrl :=  e.Request.AbsoluteURL(e.ChildAttr("[rel='next']", "href"))
		if len(nextPageUrl) == 0 {
			return
		}
		time.Sleep(REQUEST_DELAY)
		e.Request.Visit(nextPageUrl)
	})

	mainCollector.OnHTML("li[role='article']", func (e *colly.HTMLElement) {
			link:=e.Request.AbsoluteURL(e.ChildAttr("a[href^='/works/']", "href"))
			fmt.Println(link)

			//downloadDetails := new(DownloadDetails)
			
			
			//fetchSingleDownloadDetailss(dlCollector)
			dlCollector.Visit(link)
			//fmt.Println(downloadDetails)

			/*
			filePath := formatPath(downloadDetails.title, "pdf")
			downloadUrl := fmt.Sprintf("https://%s%s", DOMAIN, downloadDetails.getUrlByFormat("pdf"))
			downloadErr := downloadSingleFic(downloadUrl, filePath)
			if downloadErr != nil {
				fmt.Println(downloadErr)
				return
			}
			
	})
}

func setDlCollector() {
	dlCollector = mainCollector.Clone()
	dlCollector.OnRequest(func(req *colly.Request) {
		fmt.Printf(INFO_VISITING, req.URL.String())
	})
	dlCollector.OnHTML("h2.title.heading", titleprint)
}

func fetchAllDownloadDetailsFromSearch(ao3Url string, endPage int) {
	setmainCollector(endPage)
	setDlCollector()
	mainCollector.Visit(ao3Url)
}

func titleprint(e *colly.HTMLElement) {
	fmt.Println(e.Text)
}

func fetchSingleDownloadDetailss() {
	
	
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
			panic(fetchError) //Not sure if its a good idea to panic here but IDK man
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
*/

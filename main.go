package main

import (
	"fmt"
)

func main() {
	//create download directory
	dirCreationErr := createDirectory(DOWNLOAD_DIR)
	if dirCreationErr != nil {
		fmt.Println(dirCreationErr)
		return
	}

	//parse args
	ao3Url, format, endPage, parseErr := parseDownloadArgsTwo()
	if parseErr != nil {
		return
	}
	fmt.Println("url", ao3Url, "format", format, "endpage", endPage)
	fetchAllDownloadDetailsFromSearch(ao3Url, endPage)

	//create collector & download details struct
	//collector := createCollector()
	//downloadDetails := new(DownloadDetails)
	
	//fetchSingleDownloadDetails(ao3Url, downloadDetails)
	/*
	

	filePath := formatPath(downloadDetails.title, format)
	downloadUrl := fmt.Sprintf("https://%s%s", DOMAIN, downloadDetails.getUrlByFormat(format))
	
	downloadErr := downloadSingleFic(downloadUrl, filePath)
	if downloadErr != nil {
		fmt.Println(downloadErr)
		return
	}
	*/
}

/*
func main() {
	//create download directory
	dirCreationErr := createDirectory(DOWNLOAD_DIR)
	if dirCreationErr != nil {
		fmt.Println(dirCreationErr)
		return
	}

	//parse args
	ao3Url, format, parseErr := parseDownloadArgs()
	if parseErr != nil {
		return
	}

	//create collector & download details struct
	collector := createCollector()
	downloadDetails := new(DownloadDetails)
	
	fetchSingleDownloadDetails(collector, ao3Url, downloadDetails)
	time.Sleep(REQUEST_DELAY)

	filePath := formatPath(downloadDetails.title, format)
	downloadUrl := fmt.Sprintf("https://%s%s", DOMAIN, downloadDetails.getUrlByFormat(format))
	
	downloadErr := downloadSingleFic(downloadUrl, filePath)
	if downloadErr != nil {
		fmt.Println(downloadErr)
		return
	}
}
*/
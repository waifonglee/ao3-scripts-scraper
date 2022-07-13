package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//create download directory
	dirCreationErr := createDirectory(DOWNLOAD_DIR)
	if dirCreationErr != nil {
		fmt.Println(dirCreationErr)
		os.Exit(1)
	}

	//parse args
	ao3Url, format := parseDownloadArgs()

	//create collector & download details struct
	collector := createCollector()
	downloadDetails := new(DownloadDetails)
	
	fetchSingleDownloadDetails(collector, ao3Url, downloadDetails)
	time.Sleep(REQUEST_DELAY)

	filePath := formatPath(downloadDetails.title, format)
	downloadUrl := fmt.Sprintf("https://%s%s", DOMAIN, downloadDetails.getUrlByFormat(format))
	
	downloadErr := downloadFic(downloadUrl, filePath)
	if downloadErr != nil {
		fmt.Println(downloadErr)
		os.Exit(1)
	}
}
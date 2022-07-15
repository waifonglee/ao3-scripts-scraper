package main

import (
	"fmt"
	"strconv"
)

func main() {
	//create download directory
	dirCreationErr := createDirectory(DOWNLOAD_DIR)
	if dirCreationErr != nil {
		fmt.Println(dirCreationErr)
		return
	}

	//parse args
	values, parseErr := parseCommand()
	if parseErr != nil {
		return
	}

	fmt.Println("parse values", values)
	
	
	switch values["urlType"] {
	case "0":
		downloadWorkFromLink(values["url"], values["format"])
	case "1":
		end, _ := strconv.Atoi(values["end"])
		downloadWorksFromLink(values["url"], end, values["format"])
	}
	
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
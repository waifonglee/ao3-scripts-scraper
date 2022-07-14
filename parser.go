package main

import (
	"errors"
	"flag"
	"fmt"
)

func parseDownloadArgs() (string, string, error) {
	var url string
	var format int
	var formatStr string 
	
	flag.StringVar(&url, "url", "", "URL of fanfic to download")
	flag.IntVar(&format, "format", 0, "Format of download file\n0: pdf\n1: html\n2: mobi\n3: epub\n4: azw3\ndefault: pdf")
	flag.Parse()
	//check if url is valid e.g from page or from fic
	if len(url) == 0 || format > 4 || format < 0 {
		fmt.Println(INFO_INVALID_ARGS)
		flag.PrintDefaults()
		return "", "", errors.New(ERROR_INVALID_ARGS)
	}
	
	switch format {
	case 0:
		formatStr = "pdf"
	case 1:
		formatStr = "html"
	case 2:
		formatStr = "mobi"
	case 3:
		formatStr = "epub"
	case 4:
		formatStr = "azw"	
	}

	return url, formatStr, nil
}

func parseDownloadArgsTwo() (string, string, int, error) {
	var url string
	var format int
	var endpage int
	var formatStr string 
	
	flag.StringVar(&url, "url", "", "URL of fanfic to download")
	flag.IntVar(&format, "format", 0, "Format of download file\n0: pdf\n1: html\n2: mobi\n3: epub\n4: azw3\ndefault: pdf")
	flag.IntVar(&endpage, "endpage", 1, "Last page to download fanfic from. default: 1")
	flag.Parse()
	//check if url is valid e.g from page or from fic
	if len(url) == 0 || format > 4 || format < 0 || endpage < 0 {
		fmt.Println(endpage)
		fmt.Println(INFO_INVALID_ARGS)
		flag.PrintDefaults()
		return "", "", -1, errors.New(ERROR_INVALID_ARGS)
	}
	
	switch format {
	case 0:
		formatStr = "pdf"
	case 1:
		formatStr = "html"
	case 2:
		formatStr = "mobi"
	case 3:
		formatStr = "epub"
	case 4:
		formatStr = "azw"	
	}

	return url, formatStr, endpage, nil
}
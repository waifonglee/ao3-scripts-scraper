package main

import (
	"flag"
	"os"
)

func parseDownloaderArgs() (string, string) {
	var url string
	var format int
	var formatStr string 
	
	flag.StringVar(&url, "url", "", "URL of fanfic to download")
	flag.IntVar(&format, "format", 0, "Format of download file\n0: pdf\n1: html\n2: mobi\n3: epub\n4:azw3\ndefault: pdf")
	flag.Parse()
	
	if len(url) == 0 || format > 4 || format < 0 {
		flag.PrintDefaults()
		os.Exit(1)
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

	return url, formatStr
}
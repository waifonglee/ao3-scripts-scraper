package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func parseCommand() (map[string]string, error) {
	if len(os.Args) < 2 {
		fmt.Printf(INFO_INVALID_ARGS)
		fmt.Printf(INFO_REQUIRE_SUBCOMMAND)
		return nil, errors.New(ERROR_INVALID_ARGS)
	}

	switch os.Args[1] {
	case "download":
		return parseDownloadCommand(os.Args[2:])

	default:
		fmt.Printf(INFO_INVALID_ARGS)
		fmt.Printf(INFO_REQUIRE_SUBCOMMAND)
		return nil, errors.New(ERROR_INVALID_ARGS)
	}
}

func parseDownloadCommand(args []string) (map[string]string, error) {
	values := make(map[string]string)
	command := flag.NewFlagSet("download", flag.ExitOnError)
	urlType := command.Int("type", 0, "type of url\n0: download single fic from fic url\n1: download fics from download fics from search/bookmarks/works url url\ndefault: 0")
	url := command.String("url", "", "url to download from")
	format := command.Int("format", 0, "format of download file\n0: pdf\n1: html\n2: mobi\n3: epub\n4: azw3\ndefault: 0" )
	end := command.Int("end", 0, "last page to download\nonly applies to type 1.\ndefault: all")
	command.Parse(args)

	if len(*url) == 0 || *format > 4 || *format < 0 || *urlType > 1 || *urlType < 0 || *end < 0 {
		fmt.Printf(INFO_INVALID_ARGS)
		command.PrintDefaults()
		return nil, errors.New(ERROR_INVALID_ARGS)
	}
	
	values["urlType"] = strconv.Itoa(*urlType)
	values["url"] = *url
	values["format"] = convertFormatIntToString(*format)
	values["end"] = strconv.Itoa(*end)
	
	return values, nil

}


func convertFormatIntToString(format int) string {
	switch format {
	case 0:
		return "pdf"
	case 1:
		return "html"
	case 2:
		return "mobi"
	case 3:
		return "epub"
	case 4:
		return "azw"
	}

	return ""
}

/*
func parseDownloadArgs() (string, string, error) {
	var url string
	var format int
	var formatStr string

	flag.StringVar(&url, "url", "", "URL of fanfic to download")
	flag.IntVar(&format, "format", 0, "Format of download file\n0: pdf\n1: html\n2: mobi\n3: epub\n4: azw3\ndefault: pdf")
	flag.Parse()
	//check if url is valid e.g from page or from fic
	if len(url) == 0 || format > 4 || format < 0 {
		fmt.Printf(INFO_INVALID_ARGS)
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
		fmt.Printf(INFO_INVALID_ARGS)
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
*/
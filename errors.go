package main

import "fmt"

type FetchError struct {
	url        string
	statusCode int
	err        error
}

func (e *FetchError) Error() string {
	return fmt.Sprintf(ERROR_FETCH, e.url, e.statusCode, e.err)
}

type DownloadError struct {
	step string
	url string
	dir string
	err error
}

func (e *DownloadError) Error() string {
	return fmt.Sprintf(ERROR_DOWNLOAD, e.step, e.dir, e.url, e.err)
}

package main

type DownloadDetails struct {
	title string
	urls map[string]string
}

// fileType must be in upper case
func (download *DownloadDetails) getUrlByFormat(format string) string {
	return download.urls[format]
}





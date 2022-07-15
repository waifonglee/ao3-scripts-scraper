package main

type DownloadDetails struct {
	title string
	format string
	downloadLink string
}

func (d *DownloadDetails) resetTitleLink() {
	d.title = ""
	d.downloadLink = ""
}

type PageDetails struct {
	endPage int
}


/*
func (download *DownloadDetails) getUrlByFormat(format string) string {
	return download.urls[format]
}
*/




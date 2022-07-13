package main


const (
	INFO_VISITING = "Visiting %s\n"
	INFO_TOO_MANY_REQ = "Too many requests. Going to sleep."

	//url, statuscode, err
	ERROR_FETCH = "Fetch error: request for %s failed with status %d and error: %v\n"
	//case, dir, url, err
	ERROR_DOWNLOAD = "Download error: %s for %s from %s failed with error: %v\n"
	ERROR_DIR_CREATION = "Error: download directory creation failed with error: %v\n"
	DOWNLOAD_DIR = "downloads"
)
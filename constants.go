package main


const (
	//time, url
	INFO_VISITING = "%s: Visiting %s\n"
	INFO_TOO_MANY_REQ = "%s: Too many requests. Going to sleep.\n"
	INFO_INVALID_ARGS = "Invalid arguments. Please refer to usage below.\n"
	INFO_DOWNLOADING = "%s: Starting download from %s\n"

	ERROR_INVALID_ARGS = "Invalid arguments\n"
	//url, statuscode, err
	ERROR_FETCH = "Fetch error: request for %s failed with status %d and error: %v\n"
	//case, path, url, err
	ERROR_DOWNLOAD = "Download error: %s for %s from %s failed with error: %v\n"
	ERROR_DIR_CREATION = "Error: download directory creation failed with error: %v\n"

)
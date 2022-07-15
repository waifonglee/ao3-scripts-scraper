package main

import "time"


const (
	REQUEST_DELAY = 5 * time.Second
	RETRY_TIMEOUT = 300 * time.Second
	DOWNLOAD_DIR = "downloads"
	CACHE_DIR = "cache"
	DOMAIN = "archiveofourown.org"
)

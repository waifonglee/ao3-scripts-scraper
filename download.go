package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func downloadSingleWork(d *DownloadDetails) error {
	path := formatPath(d.title, d.format)
	fic, err := os.Create(path)
	if err != nil {
		downloadErr := &DownloadError{step: "file creation", url: d.downloadLink, path: path, err: err}
		return downloadErr
	}
	defer fic.Close()

	resp, err := http.Get(d.downloadLink)
	if err != nil {
		downloadErr := &DownloadError{step: "file fetch", url: d.downloadLink, path: path, err: err}
		return downloadErr
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusTooManyRequests {
			fmt.Printf(INFO_TOO_MANY_REQ, time.Now().String())
			time.Sleep(RETRY_TIMEOUT)
			return downloadSingleWork(d)
		}
		downloadErr := &DownloadError{step: "file fetch", url: d.downloadLink, path: path, err: errors.New(fmt.Sprintf("status : %d", resp.StatusCode))}
		return downloadErr
	}
	
	_, err = io.Copy(fic, resp.Body)
	if err != nil {
		downloadErr := &DownloadError{step: "file write", url: d.downloadLink, path: path, err: err}
		return downloadErr
	}
	
	return nil
}
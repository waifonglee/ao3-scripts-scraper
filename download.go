package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)





func downloadSingleFic(d *DownloadDetails) error {
	path := formatPath(d.title, d.format)
	fic, err := os.Create(path)
	if err != nil {
		downloadErr := &DownloadError{step: "file creation", url: d.link, path: path, err: err}
		return downloadErr
	}
	defer fic.Close()

	resp, err := http.Get(d.link)
	if err != nil {
		downloadErr := &DownloadError{step: "file fetch", url: d.link, path: path, err: err}
		return downloadErr
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		downloadErr := &DownloadError{step: "file fetch", url: d.link, path: path, err: errors.New(fmt.Sprintf("status : %d", resp.StatusCode))}
		return downloadErr
	}
	
	_, err = io.Copy(fic, resp.Body)
	if err != nil {
		downloadErr := &DownloadError{step: "file write", url: d.link, path: path, err: err}
		return downloadErr
	}
	
	return nil
}
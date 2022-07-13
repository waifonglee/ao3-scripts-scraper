package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadFic(url string, path string) error {
	fic, err := os.Create(path)
	if err != nil {
		downloadErr := &DownloadError{step: "file creation", url: url, path: path, err: err}
		return downloadErr
	}
	defer fic.Close()

	resp, err := http.Get(url)
	if err != nil {
		downloadErr := &DownloadError{step: "file fetch", url: url, path: path, err: err}
		return downloadErr
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		downloadErr := &DownloadError{step: "file fetch", url: url, path: path, err: errors.New(fmt.Sprintf("status : %d", resp.StatusCode))}
		return downloadErr
	}
	
	_, err = io.Copy(fic, resp.Body)
	if err != nil {
		downloadErr := &DownloadError{step: "file write", url: url, path: path, err: err}
		return downloadErr
	}
	
	return nil
}
package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func replaceSpaceWithUnderscore(s string) string {
	return strings.ReplaceAll(s, " ", "_")
}

func createDirectory(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return errors.New(fmt.Sprintf(ERROR_DIR_CREATION, err))
	}
	return nil
}

func formatPath(title string, format string) string {
	fileName := fmt.Sprintf("%s.%s", replaceSpaceWithUnderscore(title), format)
	return filepath.Join(DOWNLOAD_DIR, fileName)
}

func getFormattedTime() string{
	return time.Now().Format("2006-01-02 15:04:05")
}
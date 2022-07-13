package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
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
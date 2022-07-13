package main

import "strings"

func replaceSpaceWithUnderscore(s string) string {
	return strings.ReplaceAll(s, " ", "_")
}
package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	code := os.Getenv("MVP_CREATOR_ID")
	if len(code) == 0 {
		code = "WT.mc_id=AZ-MVP-5002618"
	}

	text, _ := clipboard.ReadAll()
	if isValidURL(text) && !strings.Contains(text, code) {
		querystring := "?"
		if strings.Contains(text, querystring) {
			querystring = "&"
		}
		newURL := text + querystring + code
		fmt.Println(newURL)
		clipboard.WriteAll(newURL)
	}
}

// isValidURL tests a string to determine if it is a well-structured url or not.
func isValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
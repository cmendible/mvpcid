package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/MakeNowJust/hotkey"
	"github.com/atotto/clipboard"
)

func main() {
	creatorID := os.Getenv("MVP_CREATOR_ID")
	if len(creatorID) == 0 {
		creatorID = "AZ-MVP-5002618"
	}

	code := fmt.Sprintf("WT.mc_id=%s", creatorID)

	hkey := hotkey.New()

	quit := make(chan bool)

	_, _ = hkey.Register(hotkey.Ctrl+hotkey.Shift, 'Q', func() {
		quit <- true
	})

	_, _ = hkey.Register(hotkey.Ctrl+hotkey.Shift, 'M', func() {
		text, _ := clipboard.ReadAll()
		if isValidURL(text) && !strings.Contains(text, code) {
			querystring := "?"
			if strings.Contains(text, querystring) {
				querystring = "&"
			}
			newURL := text + querystring + code
			fmt.Println(newURL)
			err := clipboard.WriteAll(newURL)
			if err != nil {
				fmt.Println(err)
			}
		}
	})

	<-quit
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

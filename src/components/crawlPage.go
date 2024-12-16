package components

import (
	"fmt"
	"net/url"
)

func CrawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
		return
	}

	// skip other websites
	if currentURL.Hostname() != baseURL.Hostname() {
		return
	}

	normalisedURL, err := NormaliseUrl(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalisedURL: %v", err)
		return
	}

	if _, visited := pages[normalisedURL]; visited {
		pages[normalisedURL]++
		return
	}

	pages[normalisedURL] = 1

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := GetHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

	nextURLs, err := GetURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v", err)
		return
	}

	for _, nextURL := range nextURLs {
		CrawlPage(rawBaseURL, nextURL, pages)
	}
}

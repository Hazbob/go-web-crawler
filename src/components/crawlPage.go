package components

import (
	"fmt"
	"net/url"
)

func (cfg *Config) CrawlPage(rawCurrentURL string) {
	cfg.ConcurrencyControl <- struct{}{}
	defer func() {
		<-cfg.ConcurrencyControl
		cfg.Wg.Done()
	}()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	// skip other websites
	if currentURL.Hostname() != cfg.BaseURL.Hostname() {
		return
	}

	normalizedURL, err := NormaliseUrl(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
		return
	}

	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := GetHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

	nextURLs, err := GetURLsFromHTML(htmlBody, cfg.BaseURL.String())
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v", err)
		return
	}

	for _, nextURL := range nextURLs {
		cfg.Wg.Add(1)
		go cfg.CrawlPage(nextURL)
	}
}

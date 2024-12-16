package main

import (
	"fmt"
	"github.com/Hazbob/go-web-crawler/src/components"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		return
	}
	scrapeUrl := os.Args[1]
	fmt.Printf("starting crawl of: %s\n", scrapeUrl)
	var pages = make(map[string]int)
	components.CrawlPage(scrapeUrl, scrapeUrl, pages)

	for normalizedURL, count := range pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}

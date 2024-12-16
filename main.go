package main

import (
	"fmt"
	"github.com/Hazbob/go-web-crawler/src/components"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		return
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		return
	}
	rawBaseURL := os.Args[1]
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("error parsing max concurrency")
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("error parsing max pages")
	}

	cfg, err := components.Configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg.Wg.Add(1)
	go cfg.CrawlPage(rawBaseURL)
	cfg.Wg.Wait()

	for normalizedURL, count := range cfg.Pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}

	return
}

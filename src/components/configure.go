package components

import (
	"fmt"
	"net/url"
	"sync"
)

type Config struct {
	Pages              map[string]int
	BaseURL            *url.URL
	Mu                 *sync.Mutex
	ConcurrencyControl chan struct{}
	Wg                 *sync.WaitGroup
}

func (cfg *Config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.Mu.Lock()
	defer cfg.Mu.Unlock()

	if _, visited := cfg.Pages[normalizedURL]; visited {
		cfg.Pages[normalizedURL]++
		return false
	}

	cfg.Pages[normalizedURL] = 1
	return true
}

func Configure(rawBaseURL string, maxConcurrency int) (*Config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	return &Config{
		Pages:              make(map[string]int),
		BaseURL:            baseURL,
		Mu:                 &sync.Mutex{},
		ConcurrencyControl: make(chan struct{}, maxConcurrency),
		Wg:                 &sync.WaitGroup{},
	}, nil
}

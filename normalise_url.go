package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func NormaliseUrl(rawUrl string) (string, error) {
	if len(rawUrl) <= 0 {
		return "", errors.New("string cannot be empty")
	}

	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %v", rawUrl)
	}

	fullPath := parsedUrl.Host + parsedUrl.Path

	fullPath = strings.ToLower(fullPath)

	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil
	//splitStr := strings.Split(url, "/")[2:4]
	//normalisedString := strings.Join(splitStr, "/")
	//fmt.Println(normalisedString)
	//normalisedString, err := netUrl.Parse(url)
	//if err != nil {
	//	errors.New("Error parsing url")
	//}
	//
	//return normalisedString.Host + normalisedString.Path, nil
	//return normalisedString, nil
}

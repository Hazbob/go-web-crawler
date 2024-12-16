package components

import "fmt"

func PrintReport(pages map[string]int, baseUrl string) {
	var formattedVisits []formattedVisit
	for url, visits := range pages {
		newVisit := formattedVisit{
			baseUrl:     url,
			totalVisits: visits,
		}
		formattedVisits = append(formattedVisits, newVisit)
	}
	var visitsLineByLine string

	for _, visit := range formattedVisits {
		visitsLineByLine += fmt.Sprintf("Found %d internal links to %s\n", visit.totalVisits, visit.baseUrl)
	}

	fmt.Printf(`
======================
Report For %s
======================
%s
`, baseUrl, visitsLineByLine)
}

type formattedVisit struct {
	baseUrl     string
	totalVisits int
}
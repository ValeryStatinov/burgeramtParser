package parser

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gocolly/colly"
)

var (
	dateR = regexp.MustCompile(`\d+\.\d+\.\d+`)
)

type DrivingLicenceSpider struct {
	url string
}

func NewDrivingLicenceSpider() *DrivingLicenceSpider {
	return &DrivingLicenceSpider{
		url: "https://service.berlin.de/terminvereinbarung/termin/all/327537/",

		// anmeldung
		// url: "https://service.berlin.de/terminvereinbarung/termin/all/120686/",
	}
}

func (dls *DrivingLicenceSpider) Crawl() []string {
	c := colly.NewCollector()
	availableDates := make([]string, 0, 10)

	c.OnHTML("td.buchbar", func(h *colly.HTMLElement) {
		date := getDate(h)
		if date != "" {
			availableDates = append(availableDates, date)
		}
	})

	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode != http.StatusOK {
			fmt.Println("Response is not 200, got", r.StatusCode)
		}
	})

	c.Visit(dls.url)
	return availableDates
}

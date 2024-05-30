package parser

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

var (
	dateR = regexp.MustCompile(`\d+\.\d+\.\d+`)
)

var (
	ErrNoAvailableTermins = errors.New("no available termins")
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

func (dls *DrivingLicenceSpider) Crawl() ([]string, error) {
	c := colly.NewCollector()
	availableDates := make([]string, 0, 10)
	var err error = nil

	c.OnHTML("td.buchbar", func(h *colly.HTMLElement) {
		date := getDate(h)
		if date != "" {
			availableDates = append(availableDates, date)
		}
	})
	c.OnHTML("h1", func(h *colly.HTMLElement) {
		if strings.Contains(h.Text, "Leider sind aktuell keine Termine") {
			err = ErrNoAvailableTermins
		}
	})

	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode != http.StatusOK {
			err = fmt.Errorf("response is not 200, got %d", r.StatusCode)
		}
	})

	c.Visit(dls.url)
	return availableDates, err
}

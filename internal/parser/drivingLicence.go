package parser

import (
	"errors"
	"fmt"
	"net/http"
	"os"
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
	var resultErr error = nil

	c.OnHTML("td.buchbar", func(h *colly.HTMLElement) {
		date := getDate(h)
		if date != "" {
			availableDates = append(availableDates, date)
		}
	})
	c.OnHTML("h1", func(h *colly.HTMLElement) {
		fmt.Println("found h1 with text", h.Text)
		if strings.Contains(h.Text, "Leider sind aktuell keine Termine") {
			resultErr = ErrNoAvailableTermins
		}
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("received response", r.StatusCode)
		f, err := os.OpenFile("./a.html", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			resultErr = err
		}
		defer f.Close()
		_, e := f.Write(r.Body)
		fmt.Println(e)

		if r.StatusCode != http.StatusOK {
			resultErr = fmt.Errorf("response is not 200, got %d", r.StatusCode)
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("error occured:", err.Error())
		fmt.Println(r.StatusCode, string(r.Body))
	})

	fmt.Println("visiting", dls.url)
	c.Visit(dls.url)
	return availableDates, resultErr
}

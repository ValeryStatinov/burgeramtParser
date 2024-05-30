package parser

import "testing"

func TestDrivingLicence(t *testing.T) {
	s := NewDrivingLicenceSpider()

	t.Run("crawl", func(t *testing.T) {
		s.Crawl()
	})
}

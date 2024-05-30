package parser

import "github.com/gocolly/colly"

func regexpDate(s string) string {
	dateMatch := dateR.FindAllString(s, -1)
	if len(dateMatch) > 0 {
		date := dateMatch[0]
		return date
	}
	return ""
}

func getDate(h *colly.HTMLElement) string {
	title := h.ChildAttr("a", "title")

	return regexpDate(title)
}

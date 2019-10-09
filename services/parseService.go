package services

import (
	"strings"

	"github.com/gocolly/colly"
)

// SearchPhrase - search for phrase on the site
func SearchPhrase(url, phrase string) bool {
	c := colly.NewCollector()
	result := false

	c.OnHTML("html", func(e *colly.HTMLElement) {
		result = strings.Contains(string(e.Response.Body), phrase)
	})

	c.Visit("http://" + url)

	return result
}

// SearchMetatag - search for metatag on the site
func SearchMetatag(url, name, content string) map[string]bool {
	c := colly.NewCollector()
	result := map[string]bool{
		"name":    false,
		"content": false,
	}

	c.OnHTML("meta[name]", func(e *colly.HTMLElement) {
		n := e.Attr("name")
		co := e.Attr("content")

		if name == n {
			result["name"] = true
			if content == co {
				result["content"] = true
			}
		}
	})

	c.Visit("http://" + url)

	return result
}

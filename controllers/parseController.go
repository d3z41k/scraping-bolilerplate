package controllers

import (
	"net/http"
	"strings"

	u "github.com/d3z41k/scraping-boilerplate/utils"
	"github.com/gocolly/colly"

	"github.com/go-chi/chi"
)

// SearchPhrase - search for phrase on the site
var SearchPhrase = func(w http.ResponseWriter, r *http.Request) {
	url := chi.URLParam(r, "url")
	phrase := chi.URLParam(r, "phrase")
	result := false

	c := colly.NewCollector()

	c.OnHTML("html", func(e *colly.HTMLElement) {
		result = strings.Contains(string(e.Response.Body), phrase)
	})

	c.Visit("http://" + url)

	if result == false {
		u.Respond(w, u.Message(404, false, "Not found"))
		return
	}

	resp := u.Message(200, result, "Ok")
	resp["url"] = url
	u.Respond(w, resp)
}

// SearchMetatag - search for metatag on the site
var SearchMetatag = func(w http.ResponseWriter, r *http.Request) {
	url := chi.URLParam(r, "url")
	name := chi.URLParam(r, "name")
	content := chi.URLParam(r, "content")
	result := map[string]bool{
		"name":    false,
		"content": false,
	}

	c := colly.NewCollector()

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

	if result["name"] == false {
		u.Respond(w, u.Message(404, false, "Metatag is not found"))
		return
	}
	if result["content"] == false {
		u.Respond(w, u.Message(404, false, "Invalid metatag content"))
		return
	}

	resp := u.Message(200, true, "Ok")
	resp["url"] = url
	u.Respond(w, resp)
}

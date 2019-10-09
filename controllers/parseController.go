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

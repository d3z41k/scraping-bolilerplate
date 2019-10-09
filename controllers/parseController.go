package controllers

import (
	"net/http"

	"github.com/d3z41k/scraping-boilerplate/services"
	u "github.com/d3z41k/scraping-boilerplate/utils"

	"github.com/go-chi/chi"
)

// SearchPhrase - get result of search for phrase on the site
var SearchPhrase = func(w http.ResponseWriter, r *http.Request) {
	url := chi.URLParam(r, "url")
	phrase := chi.URLParam(r, "phrase")

	result := services.SearchPhrase(url, phrase)

	if result == false {
		u.Respond(w, u.Message(404, false, "Not found"))
		return
	}

	resp := u.Message(200, result, "Ok")
	resp["url"] = url
	u.Respond(w, resp)
}

// SearchMetatag - get result of search for metatag on the site
var SearchMetatag = func(w http.ResponseWriter, r *http.Request) {
	url := chi.URLParam(r, "url")
	name := chi.URLParam(r, "name")
	content := chi.URLParam(r, "content")

	result := services.SearchMetatag(url, name, content)

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

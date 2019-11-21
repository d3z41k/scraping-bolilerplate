package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/d3z41k/scraping-boilerplate/models"
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
		u.Respond(w, u.Message(404, false, "Search phrase not found"))
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
		u.Respond(w, u.Message(404, false, "Metatag not found"))
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

// SearchData - get result of search for data
var SearchData = func(w http.ResponseWriter, r *http.Request) {
	var searchData []models.SearchData
	var resultData = []models.ResultData{}

	err := json.NewDecoder(r.Body).Decode(&searchData)
	if err != nil {
		u.Respond(w, u.Message(400, false, "Invalid request"))
		return
	}

	for _, data := range searchData {

		var rd = models.ResultData{}
		var rt = map[string]bool{}

		rd.URL = data.URL
		rt = services.SearchMetatag(data.URL, data.Name, data.Content)
		rt["phrase"] = services.SearchPhrase(data.URL, data.Phrase)
		s := true

		for item, status := range rt {
			if status == false {
				s = false
			}
			m := u.RsultMessage(item, status)

			rd.Results = append(rd.Results, u.Result(item, status, m))
		}
		rd.Status = s
		resultData = append(resultData, rd)
	}

	resp := u.Message(200, true, "Ok")
	resp["result"] = resultData
	u.Respond(w, resp)
}

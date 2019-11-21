package v1

import (
	"net/http"

	"github.com/d3z41k/scraping-boilerplate/controllers"

	"github.com/go-chi/chi"
)

// NewRouter returns an HTTP handler that implements the routes for the API
func NewRouter() http.Handler {
	router := chi.NewRouter()

	// Register the API routes
	router.Get("/search-phrase/{url}/{phrase}", controllers.SearchPhrase)
	router.Get("/search-tag/{url}/{name}/{content}", controllers.SearchMetatag)
	router.Post("/search-data", controllers.SearchData)
	return router
}

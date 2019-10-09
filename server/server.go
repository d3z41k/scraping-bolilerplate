package server

import (
	"net/http"
	"time"

	v1 "github.com/d3z41k/scraping-boilerplate/api/v1"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// NewRouter return HTTP handler that implements the main server routers
func NewRouter() http.Handler {
	router := chi.NewRouter()

	// Set up our middleware with sane defaults
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// Set up API
	router.Mount("/api/v1", v1.NewRouter())

	return router
}

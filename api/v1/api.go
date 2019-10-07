package v1

import (
	"net/http"

	"github.com/go-chi/chi"
)

// HelloWorld returns a basic "Hello World!" message
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

// NewRouter returns an HTTP handler that implements the routes for the API
func NewRouter() http.Handler {
	router := chi.NewRouter()

	// Register the API routes)
	router.Get("/hello", HelloWorld)

	return router
}

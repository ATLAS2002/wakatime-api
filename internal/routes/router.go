package routes

import (
	"net/http"

	"github.com/ATLAS2002/wakatime-api/internal/v1/handlers"
)

// GetRouter returns a new http.ServeMux with the routes for the API.
func GetV1Router() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /{username}/coding-time", handlers.HandleCodingTime)
  router.HandleFunc("GET /{username}/languages", handlers.HandleTopLanguages)
  router.HandleFunc("GET /{username}/language/{lang}", handlers.HandleLanguage)

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	return v1
}

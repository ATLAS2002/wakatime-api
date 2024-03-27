package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ATLAS2002/wakatime-api/internal/v1/services"
)

func HandleLanguage(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
  language := r.PathValue("lang")
  log.Printf("%s asks for language: %s\n", username, language)

	stats, err := services.StatsService(username)
	if err != nil {
		http.Error(w, "Error getting stats", http.StatusBadRequest)
		panic(err)
	}

	res, err := stats.GetLanguage(language)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
  }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
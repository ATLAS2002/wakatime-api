package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ATLAS2002/wakatime-api/internal/v1/services"
)

func HandleOS(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
  os := r.PathValue("os")
  log.Printf("%s asks for os: %s\n", username, os)

	stats, err := services.StatsService(username)
	if err != nil {
		http.Error(w, "Error getting stats", http.StatusBadRequest)
		panic(err)
	}

	res, err := stats.GetOS(os)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
  }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
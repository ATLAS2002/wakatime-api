package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ATLAS2002/wakatime-api/internal/v1/services"
)

func HandleTopOS(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
  log.Printf("%s asks for top operating systems\n", username)

  limitParam := r.URL.Query().Get("limit")
  limit, err := strconv.Atoi(limitParam)
  if err != nil {
    limit = -1
  }

	stats, err := services.StatsService(username)
	if err != nil {
		http.Error(w, "Error getting stats", http.StatusBadRequest)
	}

	res := stats.GetTopOS(limit)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
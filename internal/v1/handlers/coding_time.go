package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ATLAS2002/wakatime-api/internal/v1/services"
)

func HandleCodingTime(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
  log.Printf("%s asks for coding time\n", username)

	stats, err := services.StatsService(username)
	if err != nil {
		http.Error(w, "Error getting stats", http.StatusBadRequest)
		panic(err)
	}

	totalHours, totalMinutes := stats.GetCodingTime()
	res := map[string]int {
		"hours":   totalHours,
		"minutes": totalMinutes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
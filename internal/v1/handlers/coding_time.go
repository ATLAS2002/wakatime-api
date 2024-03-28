package handlers

import (
	"net/http"

	"github.com/ATLAS2002/wakatime-api/internal/v1/middlewares"
	"github.com/ATLAS2002/wakatime-api/internal/v1/services"
	"github.com/ATLAS2002/wakatime-api/internal/v1/utils"
)

func HandleCodingTime(w http.ResponseWriter, r *http.Request) {
	stats := r.Context().Value(middlewares.StatsKey).(*services.StatsSchema)

	totalHours, totalMinutes := stats.GetCodingTime()
	res := map[string]int{
		"hours":   totalHours,
		"minutes": totalMinutes,
	}

	utils.SendResponse(w, http.StatusOK, res)
}

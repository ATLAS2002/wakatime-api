package handlers

import (
	"net/http"

	"github.com/ATLAS2002/wakatime-api/internal/v1/middlewares"
	"github.com/ATLAS2002/wakatime-api/internal/v1/services"
	"github.com/ATLAS2002/wakatime-api/internal/v1/utils"
)

func HandleTopOS(w http.ResponseWriter, r *http.Request) {
	stats := r.Context().Value(middlewares.StatsKey).(*services.StatsSchema)
	limit := utils.ParseLimit(r)

	res := stats.GetTopOS(limit)

	utils.SendResponse(w, http.StatusOK, res)
}

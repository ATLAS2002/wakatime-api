package handlers

import (
	"net/http"

	"github.com/ATLAS2002/wakatime-api/internal/v1/middlewares"
	"github.com/ATLAS2002/wakatime-api/internal/v1/services"
	"github.com/ATLAS2002/wakatime-api/internal/v1/utils"
)

func HandleOS(w http.ResponseWriter, r *http.Request) {
	stats := r.Context().Value(middlewares.StatsKey).(*services.StatsSchema)
	os := r.PathValue("os")

	res, err := stats.GetOS(os)
	if err != nil {
		utils.SendResponse(w, err.Code, err)
		return
	}

	utils.SendResponse(w, http.StatusOK, res)
}

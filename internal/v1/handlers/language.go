package handlers

import (
	"net/http"

	"github.com/ATLAS2002/wakatime-api/internal/v1/middlewares"
	"github.com/ATLAS2002/wakatime-api/internal/v1/services"
	"github.com/ATLAS2002/wakatime-api/internal/v1/utils"
)

func HandleLanguage(w http.ResponseWriter, r *http.Request) {
	stats := r.Context().Value(middlewares.StatsKey).(*services.StatsSchema)
	language := r.PathValue("lang")

	res, err := stats.GetLanguage(language)
	if err != nil {
		utils.SendResponse(w, err.Code, err)
		return
	}

	utils.SendResponse(w, http.StatusOK, res)
}

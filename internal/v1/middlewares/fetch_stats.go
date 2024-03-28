package middlewares

import (
	"context"
	"net/http"

	"github.com/ATLAS2002/wakatime-api/internal/v1/services"
	"github.com/ATLAS2002/wakatime-api/internal/v1/utils"
)

type ContextKey string

const StatsKey ContextKey = "stats"

func FetchStats(next http.HandlerFunc) http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    username := r.PathValue("username")

    stats, err := services.StatsService(username)
    if err != nil {
      utils.SendResponse(w, err.Code, err)
      return
    }

    ctx := context.WithValue(r.Context(), StatsKey, stats)
    next.ServeHTTP(w, r.WithContext(ctx))
  })
}
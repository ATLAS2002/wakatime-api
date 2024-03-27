package services

import (
	"encoding/json"
	"io"
	"net/http"
)

const baseURL = "https://wakatime.com/api/v1/users/"

// StatsService fetches stats from the WakaTime API.
func StatsService(username string) (*StatsSchema, error) {
	endpoint := baseURL + username + "/stats?is_including_today=true"
	res, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var statsData StatsSchema
	if err = json.Unmarshal(body, &statsData); err != nil {
		return nil, err
	}

	return &statsData, err
}

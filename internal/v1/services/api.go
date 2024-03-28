package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ATLAS2002/wakatime-api/internal/v1/utils"
)

const baseURL = "https://wakatime.com/api/v1/users/"

// StatsService fetches stats from the WakaTime API.
func StatsService(username string) (*StatsSchema, *utils.Error) {
	endpoint := baseURL + username + "/stats?is_including_today=true"
	res, err := http.Get(endpoint)

	if res.StatusCode == http.StatusNotFound {
		return nil, &utils.Error{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("User %s not found", username),
		}
	} else if err != nil {
		return nil, utils.DefaultError(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, utils.DefaultError(err)
	}

	var statsData StatsSchema
	if err = json.Unmarshal(body, &statsData); err != nil {
		return nil, utils.DefaultError(err)
	}

	return &statsData, nil
}

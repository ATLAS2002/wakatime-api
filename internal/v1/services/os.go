package services

import (
	"fmt"

	"github.com/ATLAS2002/wakatime-api/internal/v1/utils"
)

// GetTopOS returns top operating systems for a user.
func (stats StatsSchema) GetTopOS(limit int) []WithPercentage {
	if limit == -1 || limit > len(stats.Data.OperatingSystems) {
		return stats.Data.OperatingSystems
	} else {
		return stats.Data.OperatingSystems[:limit]
	}
}

// GetOS returns data of how much time a user spent on a specific operating system.
func (stats StatsSchema) GetOS(os string) (*WithPercentageAndRank, *utils.Error) {
	for idx, operatingSystem := range stats.Data.OperatingSystems {
		if operatingSystem.Name == os {
			return &WithPercentageAndRank{
				WithPercentage: operatingSystem,
				Rank:           idx + 1,
			}, nil
		}
	}

	return nil, &utils.Error{
		Code:    404,
		Message: fmt.Sprintf("OS %s not found", os),
	}
}

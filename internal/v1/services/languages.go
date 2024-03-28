package services

import (
	"fmt"

	"github.com/ATLAS2002/wakatime-api/internal/v1/utils"
)

// GetTopLanguages returns top languages for a user.
func (stats StatsSchema) GetTopLanguages(limit int) []WithPercentage {
	if limit == -1 || limit > len(stats.Data.Languages) {
		return stats.Data.Languages
	} else {
		return stats.Data.Languages[:limit]
	}
}

// GetLanguage returns data of how much time a user spent on a specific language.
func (stats StatsSchema) GetLanguage(lang string) (*WithPercentageAndRank, *utils.Error) {
	for idx, language := range stats.Data.Languages {
		if language.Name == lang {
			return &WithPercentageAndRank{
				WithPercentage: language,
				Rank:           idx + 1,
			}, nil
		}
	}

	return nil, &utils.Error{
		Code:    404,
		Message: fmt.Sprintf("Language %s not found", lang),
	}
}

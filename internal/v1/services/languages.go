package services

import "fmt"

// GetTopLanguages returns top languages for a user.
func (stats StatsSchema) GetTopLanguages(limit int) []WithPercentage {
  if limit == -1 || limit > len(stats.Data.Languages) {
	  return stats.Data.Languages
  } else {
    return stats.Data.Languages[:limit]    
  }
}

// GetLanguage returns data of how much time a user spent on a specific language.
func (stats StatsSchema) GetLanguage(lang string) (*WithPercentageAndRank, error) {
  for idx, language := range stats.Data.Languages {
    if language.Name == lang {
      return &WithPercentageAndRank{
        WithPercentage: language,
        Rank: idx + 1,
      }, nil
    }
  }
  return nil, fmt.Errorf("language %s not found", lang)
}
package services

// GetCodingTime returns the total coding time for a user.
func (stats *StatsSchema) GetCodingTime() (int, int) {
  var totalHours, totalMinutes int

  for _, category := range stats.Data.Categories {
    if category.Name == "Coding" {
      totalHours = category.Hours
      totalMinutes = category.Minutes
      break
    }
  }

  return totalHours, totalMinutes
}
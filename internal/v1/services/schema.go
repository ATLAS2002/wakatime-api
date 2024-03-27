package services

type StatsSchema struct {
	Data struct {
		Categories []Category `json:"categories"`
    Languages []Language `json:"languages"`
	} `json:"data"`
}

type Category struct {
	Name  string `json:"name"`
	Hours int 	 `json:"hours"`
	Minutes int  `json:"minutes"`
}

type Language struct {
  Name  string `json:"name"`
  Hours int   `json:"hours"`
  Minutes int `json:"minutes"`
  Percent float64 `json:"percent"`
}

package services

type StatsSchema struct {
	Data struct {
		Categories []BasicInfo `json:"categories"`
    Languages []WithPercentage `json:"languages"`
    Editors []WithPercentage `json:"editors"`
    OperatingSystems []WithPercentage `json:"operating_systems"`
	} `json:"data"`
}

type BasicInfo struct {
	Name  string `json:"name"`
	Hours int 	 `json:"hours"`
	Minutes int  `json:"minutes"`
}

type WithPercentage struct {
  BasicInfo
  Percent float64 `json:"percent"`
}

type WithPercentageAndRank struct {
  WithPercentage
  Rank int `json:"rank"`  
}

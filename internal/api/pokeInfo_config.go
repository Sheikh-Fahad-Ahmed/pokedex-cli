package api

type Pokemon struct {
	Name string `json:"name"`
	BaseEXP int `json:"base_experience"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []Stats
}

type Stats struct {
	BaseStat int `json:"base_stat"`
	StatInfo StatInfo `json:"Stat"`
}

type StatInfo struct {
	Name string `json:"name"`
}

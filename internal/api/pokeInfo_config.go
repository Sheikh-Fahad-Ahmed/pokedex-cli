package api

type Pokemon struct {
	Name    string  `json:"name"`
	BaseEXP int     `json:"base_experience"`
	Height  int     `json:"height"`
	Weight  int     `json:"weight"`
	Stats   []Stats `json:"stats"`
	Types   []Types `json:"types"`
}

type Stats struct {
	BaseStat int      `json:"base_stat"`
	StatInfo StatInfo `json:"Stat"`
}

type StatInfo struct {
	Name string `json:"name"`
}

type Types struct {
	Type Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
}

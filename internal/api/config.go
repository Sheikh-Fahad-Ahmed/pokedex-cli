package api

type Config struct {
	Count    int     `json:"-"`
	Results  []Item  `json:"results"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}

type Item struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

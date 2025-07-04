package api

type PokemonEncountersResponse struct {
	PokemonEncounters []EncounterConfig `json:"pokemon_encounters"`
}

type EncounterConfig struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

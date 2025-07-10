package api

type PokemonEncountersResponse struct {
	PokemonEncountersConfig []EncounterConfig `json:"pokemon_encounters"`
}

type EncounterConfig struct {
	PokemonEncounter PokemonEncounter `json:"pokemon"`
}

type PokemonEncounter struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

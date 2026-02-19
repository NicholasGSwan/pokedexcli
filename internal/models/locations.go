package models

// type LocationArea struct {
// 	Id                   int
// 	Name                 string
// 	GameIndex            int
// 	EncounterMethodRates []EncounterMethodRate
// 	Location             Location
// 	Names                []Name
// 	PokemonEncounters    []PokemonEncounter
// }

type ShortLocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreaGetResult struct {
	Count    int                 `json:"count"`
	Next     string              `json:"next"`
	Previous string              `json:"previous"`
	Results  []ShortLocationArea `json:"results"`
}

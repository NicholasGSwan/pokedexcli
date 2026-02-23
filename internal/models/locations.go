package models

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

package models

type SearchResult struct {
	Artists   []Artist      `json:"artists"`
	Locations []Locationsss `json:"locations"`
}

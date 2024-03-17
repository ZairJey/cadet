package models

type Relations struct {
	Index []Relationss `json:"index"`
}
type Relationss struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

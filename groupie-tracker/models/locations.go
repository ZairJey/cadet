package models

type Locations struct {
	Index []Locationss `json:"index"`
}

type Locationss struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Locationsss struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Locations []string `json:"locations"`
}

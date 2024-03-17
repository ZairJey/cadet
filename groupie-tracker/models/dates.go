package models

type Dates struct {
	Index []Datess `json:"index"`
}
type Datess struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

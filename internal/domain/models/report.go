package models

type Report struct {
	Statistics Statistics `json:"statistics"`
	Incidents  []Incident `json:"incidents"`
}

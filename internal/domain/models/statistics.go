package models

type Statistics struct {
	TotalValid     int            `json:"total_valid"`
	TotalMalformed int            `json:"total_malformed"`
	LevelCounts    map[string]int `json:"level_counts"`
	ServiceCounts  map[string]int `json:"service_counts"`
}

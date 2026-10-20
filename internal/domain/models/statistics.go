package models

type Statistics struct {
	TotalValid     int
	TotalMalformed int
	LevelCounts    map[string]int
	ServiceCounts  map[string]int
}

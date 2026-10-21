package models

import (
	"time"

	"github.com/Manizmn84/GhostOps/internal/domain/enums"
)

type Incident struct {
	Type      string         `json:"type"`
	Severity  enums.Severity `json:"severity"`
	Service   string         `json:"service"`
	Count     int            `json:"count"`
	StartTime time.Time      `json:"start_time"`
	EndTime   time.Time      `json:"end_time"`
}

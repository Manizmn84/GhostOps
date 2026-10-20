package models

import (
	"time"

	"github.com/Manizmn84/GhostOps/internal/domain/enums"
)

type EntryLog struct {
	Time          time.Time
	SeverityLevel enums.Severity
	ServiceName   string
	Message       string
}

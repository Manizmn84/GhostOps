package incidentrule

import (
	"github.com/Manizmn84/GhostOps/internal/domain/enums"
	"github.com/Manizmn84/GhostOps/internal/domain/models"
)

type FatalCrashRule struct{}

func (r FatalCrashRule) Match(entry models.EntryLog) (models.IncidentMatch, bool) {
	if entry.SeverityLevel != enums.Fatal {
		return models.IncidentMatch{}, false
	}

	return models.IncidentMatch{
		Type:     "FATAL_CRASH",
		Severity: enums.Fatal,
	}, true
}

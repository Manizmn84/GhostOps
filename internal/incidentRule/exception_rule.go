package incidentrule

import (
	"strings"

	"github.com/Manizmn84/GhostOps/internal/domain/enums"
	"github.com/Manizmn84/GhostOps/internal/domain/models"
)

type ApplicationExceptionRule struct{}

func (r ApplicationExceptionRule) Match(entry models.EntryLog) (models.IncidentMatch, bool) {
	if entry.SeverityLevel != enums.Error && entry.SeverityLevel != enums.Fatal {
		return models.IncidentMatch{}, false
	}

	if !strings.Contains(strings.ToLower(entry.Message), "exception") {
		return models.IncidentMatch{}, false
	}

	return models.IncidentMatch{
		Type:     "APPLICATION_EXCEPTION",
		Severity: entry.SeverityLevel,
	}, true
}

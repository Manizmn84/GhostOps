package incidentrule

import (
	"strings"

	"github.com/Manizmn84/GhostOps/internal/domain/enums"
	"github.com/Manizmn84/GhostOps/internal/domain/models"
)

type DatabaseFailureRule struct{}

func (r DatabaseFailureRule) Match(entry models.EntryLog) (models.IncidentMatch, bool) {
	if entry.SeverityLevel != enums.Error && entry.SeverityLevel != enums.Fatal {
		return models.IncidentMatch{}, false
	}

	message := strings.ToLower(entry.Message)
	matched := strings.Contains(message, "database") ||
		strings.Contains(message, "sql") ||
		strings.Contains(message, "connection pool")

	if !matched {
		return models.IncidentMatch{}, false
	}

	return models.IncidentMatch{
		Type:     "DATABASE_FAILURE",
		Severity: entry.SeverityLevel,
	}, true
}

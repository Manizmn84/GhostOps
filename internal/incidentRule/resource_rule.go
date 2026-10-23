package incidentrule

import (
	"strings"

	"github.com/Manizmn84/GhostOps/internal/domain/enums"
	"github.com/Manizmn84/GhostOps/internal/domain/models"
)

type ResourcePressureRule struct{}

func (r ResourcePressureRule) Match(entry models.EntryLog) (models.IncidentMatch, bool) {
	if entry.SeverityLevel != enums.Warn {
		return models.IncidentMatch{}, false
	}

	message := strings.ToLower(entry.Message)
	matched := strings.Contains(message, "high memory") ||
		strings.Contains(message, "reaching threshold")

	if !matched {
		return models.IncidentMatch{}, false
	}

	return models.IncidentMatch{
		Type:     "RESOURCE_PRESSURE",
		Severity: enums.Warn,
	}, true
}

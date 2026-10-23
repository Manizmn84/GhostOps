package incidentrule

import (
	"strings"

	"github.com/Manizmn84/GhostOps/internal/domain/enums"
	"github.com/Manizmn84/GhostOps/internal/domain/models"
)

type DeprecatedAPIRule struct{}

func (r DeprecatedAPIRule) Match(entry models.EntryLog) (models.IncidentMatch, bool) {
	if entry.SeverityLevel != enums.Warn {
		return models.IncidentMatch{}, false
	}

	if !strings.Contains(strings.ToLower(entry.Message), "deprecated api") {
		return models.IncidentMatch{}, false
	}

	return models.IncidentMatch{
		Type:     "DEPRECATED_API",
		Severity: enums.Warn,
	}, true
}

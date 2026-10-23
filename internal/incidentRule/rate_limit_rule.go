package incidentrule

import (
	"strings"

	"github.com/Manizmn84/GhostOps/internal/domain/enums"
	"github.com/Manizmn84/GhostOps/internal/domain/models"
)

type RateLimitRule struct{}

func (r RateLimitRule) Match(entry models.EntryLog) (models.IncidentMatch, bool) {
	if entry.SeverityLevel != enums.Warn {
		return models.IncidentMatch{}, false
	}

	if !strings.Contains(strings.ToLower(entry.Message), "rate limit") {
		return models.IncidentMatch{}, false
	}

	return models.IncidentMatch{
		Type:     "RATE_LIMIT_WARNING",
		Severity: enums.Warn,
	}, true
}

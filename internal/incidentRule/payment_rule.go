package incidentrule

import (
	"strings"

	"github.com/Manizmn84/GhostOps/internal/domain/enums"
	"github.com/Manizmn84/GhostOps/internal/domain/models"
)

type PaymentFailureRule struct{}

func (r PaymentFailureRule) Match(
	entry models.EntryLog,
) (models.IncidentMatch, bool) {
	message := strings.ToLower(entry.Message)

	matched := strings.Contains(message, "payment gateway") ||
		strings.Contains(message, "bad gateway")

	if !matched {
		return models.IncidentMatch{}, false
	}

	return models.IncidentMatch{
		Type:     "PAYMENT_FAILURE",
		Severity: enums.Error,
	}, true
}

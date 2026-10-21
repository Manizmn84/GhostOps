package incidentrule

import (
	"github.com/Manizmn84/GhostOps/internal/domain/models"
)

type IncidentRule interface {
	Match(entry models.EntryLog) (models.IncidentMatch, bool)
}

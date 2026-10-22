package reporter

import (
	"io"

	"github.com/Manizmn84/GhostOps/internal/domain/models"
)

type Reporter interface {
	Write(w io.Writer, report models.Report) error
}

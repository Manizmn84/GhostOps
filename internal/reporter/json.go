package reporter

import (
	"encoding/json"
	"io"

	"github.com/Manizmn84/GhostOps/internal/domain/models"
)

type JSONReporter struct{}

func (r JSONReporter) Write(
	w io.Writer,
	report models.Report,
) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")

	return encoder.Encode(report)
}

package reporter

import (
	"fmt"
	"io"
	"sort"
	"time"

	"github.com/Manizmn84/GhostOps/internal/domain/enums"
	"github.com/Manizmn84/GhostOps/internal/domain/models"
)

type TextReporter struct{}

func (r TextReporter) Write(
	w io.Writer,
	report models.Report,
) error {
	writef := func(format string, args ...any) error {
		_, err := fmt.Fprintf(w, format, args...)
		return err
	}

	if err := writef("GhostOps Log Analysis Report\n"); err != nil {
		return err
	}
	if err := writef("============================\n\n"); err != nil {
		return err
	}
	if err := writef("Summary\n"); err != nil {
		return err
	}
	if err := writef("  Valid entries: %d\n", report.Statistics.TotalValid); err != nil {
		return err
	}
	if err := writef("  Malformed lines: %d\n\n", report.Statistics.TotalMalformed); err != nil {
		return err
	}

	if err := writef("Entries by severity\n"); err != nil {
		return err
	}
	severityOrder := []enums.Severity{
		enums.Info,
		enums.Warn,
		enums.Error,
		enums.Fatal,
	}
	for _, severity := range severityOrder {
		if err := writef(
			"  %-5s %d\n",
			severity,
			report.Statistics.LevelCounts[string(severity)],
		); err != nil {
			return err
		}
	}

	if err := writef("\nEntries by service\n"); err != nil {
		return err
	}
	services := make([]string, 0, len(report.Statistics.ServiceCounts))
	for service := range report.Statistics.ServiceCounts {
		services = append(services, service)
	}
	sort.Strings(services)

	if len(services) == 0 {
		if err := writef("  No services found.\n"); err != nil {
			return err
		}
	} else {
		for _, service := range services {
			if err := writef(
				"  %-24s %d\n",
				service,
				report.Statistics.ServiceCounts[service],
			); err != nil {
				return err
			}
		}
	}

	if err := writef("\nIncidents\n"); err != nil {
		return err
	}
	incidents := append([]models.Incident(nil), report.Incidents...)
	sort.Slice(incidents, func(i, j int) bool {
		if incidents[i].Type != incidents[j].Type {
			return incidents[i].Type < incidents[j].Type
		}
		if incidents[i].Service != incidents[j].Service {
			return incidents[i].Service < incidents[j].Service
		}
		return incidents[i].StartTime.Before(incidents[j].StartTime)
	})

	if len(incidents) == 0 {
		return writef("  No incidents detected.\n")
	}

	for index, incident := range incidents {
		if err := writef("  %d. %s\n", index+1, incident.Type); err != nil {
			return err
		}
		if err := writef("     Severity: %s\n", incident.Severity); err != nil {
			return err
		}
		if err := writef("     Service: %s\n", incident.Service); err != nil {
			return err
		}
		if err := writef("     Count: %d\n", incident.Count); err != nil {
			return err
		}
		if err := writef(
			"     Time range: %s - %s\n",
			incident.StartTime.Format(time.RFC3339),
			incident.EndTime.Format(time.RFC3339),
		); err != nil {
			return err
		}
	}

	return nil
}

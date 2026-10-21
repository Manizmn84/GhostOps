package analyzer

import (
	"github.com/Manizmn84/GhostOps/internal/domain/models"
	incidentrule "github.com/Manizmn84/GhostOps/internal/incidentRule"
)

type Analyzer struct {
	statistics models.Statistics
	incidents  map[string]*models.Incident
	rules      []incidentrule.IncidentRule
}

func NewAnalyzer(rules []incidentrule.IncidentRule) *Analyzer {
	return &Analyzer{
		statistics: models.Statistics{
			LevelCounts:   make(map[string]int),
			ServiceCounts: make(map[string]int),
		},
		incidents: make(map[string]*models.Incident),
		rules:     rules,
	}
}

func (a *Analyzer) AddMalformed() {
	a.statistics.TotalMalformed++
}

func (a *Analyzer) addIncident(
	entry models.EntryLog,
	match models.IncidentMatch,
) {
	key := match.Type + "|" + entry.ServiceName

	incident, exists := a.incidents[key]
	if !exists {
		a.incidents[key] = &models.Incident{
			Type:      match.Type,
			Severity:  match.Severity,
			Service:   entry.ServiceName,
			Count:     1,
			StartTime: entry.Time,
			EndTime:   entry.Time,
		}
		return
	}

	incident.Count++

	if entry.Time.Before(incident.StartTime) {
		incident.StartTime = entry.Time
	}

	if entry.Time.After(incident.EndTime) {
		incident.EndTime = entry.Time
	}
}

func (a *Analyzer) AddEntry(entry models.EntryLog) {
	a.statistics.TotalValid++
	a.statistics.LevelCounts[string(entry.SeverityLevel)]++
	a.statistics.ServiceCounts[entry.ServiceName]++

	for _, rule := range a.rules {
		match, ok := rule.Match(entry)
		if !ok {
			continue
		}
		a.addIncident(entry, match)

	}
}

func (a *Analyzer) Report() models.Report

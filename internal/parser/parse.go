package parser

import (
	"errors"
	"strings"
	"time"

	"github.com/Manizmn84/GhostOps/internal/domain/enums"
	"github.com/Manizmn84/GhostOps/internal/domain/models"
)

var ErrMalformedLine = errors.New("malformed log line")

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(input string) (models.EntryLog, error) {
	entryValue := strings.SplitN(strings.TrimSpace(input), " ", 4)

	if len(entryValue) != 4 {
		return models.EntryLog{}, ErrMalformedLine
	}

	timestamp, err := time.Parse(time.RFC3339, entryValue[0])
	if err != nil {
		return models.EntryLog{}, ErrMalformedLine
	}

	if !enums.IsSeverity(entryValue[1]) {
		return models.EntryLog{}, ErrMalformedLine
	}

	if strings.TrimSpace(entryValue[2]) == "" {
		return models.EntryLog{}, ErrMalformedLine
	}

	message := strings.TrimSpace(entryValue[3])
	if message == "" {
		return models.EntryLog{}, ErrMalformedLine
	}

	return models.EntryLog{
		Time:          timestamp,
		SeverityLevel: enums.Severity(entryValue[1]),
		ServiceName:   entryValue[2],
		Message:       message,
	}, nil
}

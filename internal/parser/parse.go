package parser

import (
	"errors"
	"strings"
	"time"

	"github.com/Manizmn84/GhostOps/internal/domain/enums"
	"github.com/Manizmn84/GhostOps/internal/domain/models"
)

var ValidationErr error = errors.New("Validation Error")
var MalformedLineErr error = errors.New("Malformed Line Error")

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(input string) (models.EntryLog, error) {
	entryValue := strings.SplitN(input, " ", 4)

	if len(entryValue) != 4 {
		return models.EntryLog{}, MalformedLineErr
	}

	timestamp, err := time.Parse(time.RFC1123, entryValue[0])
	if err != nil {
		return models.EntryLog{}, MalformedLineErr
	}

	if !enums.IsSeverity(entryValue[1]) {
		return models.EntryLog{}, MalformedLineErr
	}

	if strings.TrimSpace(entryValue[2]) == "" {
		return models.EntryLog{}, MalformedLineErr
	}

	return models.EntryLog{
		Time:          timestamp,
		SeverityLevel: enums.Severity(entryValue[1]),
		ServiceName:   entryValue[2],
		Message:       entryValue[3],
	}, nil
}

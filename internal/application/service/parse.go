package service

import (
	"errors"
	"strings"
	"time"

	"github.com/Manizmn84/GhostOps/internal/domain/enums"
	"github.com/Manizmn84/GhostOps/internal/domain/models"
)

var timeLayout string = "2006-01-02"

var ValidationErr error = errors.New("validation Error")

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(input string) (models.EntryLog, error) {
	entryValue := strings.SplitN(input, " ", 4)

	timestamp, err := time.Parse(timeLayout, entryValue[0])
	if err != nil {
		return models.EntryLog{}, ValidationErr
	}

	if !enums.IsSeverity(entryValue[1]) {
		return models.EntryLog{}, ValidationErr
	}

	if strings.TrimSpace(entryValue[2]) == "" {
		return models.EntryLog{}, ValidationErr
	}

	return models.EntryLog{
		Time:          timestamp,
		SeverityLevel: enums.Severity(entryValue[1]),
		ServiceName:   entryValue[2],
		Message:       entryValue[3],
	}, nil
}

package reader

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"github.com/Manizmn84/GhostOps/internal/analyzer"
	parserService "github.com/Manizmn84/GhostOps/internal/parser"
)

func Reader(
	reader io.Reader,
	parser parserService.Parser,
	analyzer analyzer.Analyzer,
) error {
	scanner := bufio.NewScanner(reader)

	scanner.Buffer(make([]byte, 64*1024), 1024*1024)

	for scanner.Scan() {
		input := scanner.Text()

		entry, err := parser.Parse(input)
		if err != nil {
			if errors.Is(err, parserService.MalformedLineErr) {
				analyzer.AddMalformed()
				continue
			}
			return err
		}
		analyzer.AddEntry(entry)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("read log file: %w", err)
	}

	return nil
}

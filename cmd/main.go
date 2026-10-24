package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Manizmn84/GhostOps/internal/analyzer"
	incidentrule "github.com/Manizmn84/GhostOps/internal/incidentRule"
	"github.com/Manizmn84/GhostOps/internal/parser"
	"github.com/Manizmn84/GhostOps/internal/reader"
	"github.com/Manizmn84/GhostOps/internal/reporter"
)

func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(args []string, stdout, stderr io.Writer) error {
	flags := flag.NewFlagSet("ghostops", flag.ContinueOnError)
	flags.SetOutput(stderr)

	outputPath := flags.String("o", "", "write the report as JSON to a file")
	flags.Usage = func() {
		fmt.Fprintln(stderr, "Usage: ghostops [-o report.json] <log-file>")
		flags.PrintDefaults()
	}

	if err := flags.Parse(args); err != nil {
		return err
	}

	if flags.NArg() != 1 {
		flags.Usage()
		return errors.New("exactly one log file must be provided")
	}

	inputFile, err := os.Open(flags.Arg(0))
	if err != nil {
		return fmt.Errorf("open input file: %w", err)
	}

	rules := []incidentrule.IncidentRule{
		incidentrule.FatalCrashRule{},
		incidentrule.DatabaseFailureRule{},
		incidentrule.PaymentFailureRule{},
		incidentrule.ApplicationExceptionRule{},
		incidentrule.ResourcePressureRule{},
		incidentrule.RateLimitRule{},
		incidentrule.DeprecatedAPIRule{},
	}

	logAnalyzer := analyzer.NewAnalyzer(rules)
	processErr := reader.Reader(inputFile, parser.NewParser(), logAnalyzer)
	closeErr := inputFile.Close()

	if processErr != nil {
		return processErr
	}
	if closeErr != nil {
		return fmt.Errorf("close input file: %w", closeErr)
	}

	report := logAnalyzer.Report()
	if *outputPath == "" {
		return (reporter.TextReporter{}).Write(stdout, report)
	}

	outputFile, err := os.Create(*outputPath)
	if err != nil {
		return fmt.Errorf("create output file: %w", err)
	}

	writeErr := (reporter.JSONReporter{}).Write(outputFile, report)
	closeErr = outputFile.Close()

	if writeErr != nil {
		return fmt.Errorf("write JSON report: %w", writeErr)
	}
	if closeErr != nil {
		return fmt.Errorf("close output file: %w", closeErr)
	}

	return nil
}

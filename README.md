# GhostOps

GhostOps is a command-line log analyzer written in Go. It reads a log file line by line, calculates statistics, detects known incident patterns, and prints a text report or writes a JSON report.

The project only uses the Go standard library.

## Log format

Each valid line must have these fields:

```text
<timestamp> <level> <service> <message>
```

Example:

```text
2026-10-02T10:22:03Z ERROR payment-service Payment gateway timeout - transaction aborted
```

The timestamp must use RFC3339 format. Supported levels are `INFO`, `WARN`, `ERROR`, and `FATAL`.

Invalid lines are skipped and counted as malformed. They do not stop the rest of the analysis.

## Build

```bash
go build -o ghostops ./cmd
```

## Usage

Print a human-readable report:

```bash
./ghostops sample.log
```

Write a JSON report to a file:

```bash
./ghostops -o report.json sample.log
```

The `-o` flag must be placed before the input file.

The program exits with a non-zero status when arguments are invalid, the input cannot be read, or the output cannot be written.

## Example output

Shortened output for the included `sample.log`:

```text
GhostOps Log Analysis Report
============================

Summary
  Valid entries: 5532
  Malformed lines: 468

Entries by severity
  INFO  3335
  WARN  1106
  ERROR 820
  FATAL 271

Entries by service
  auth-service             1106
  notification-service     1130
  order-service            1082
  payment-service          1090
  user-service             1124

Incidents
  1. APPLICATION_EXCEPTION
     Severity: ERROR
     Service: auth-service
     Count: 52
     Time range: 2026-10-01T04:27:37Z - 2026-10-06T17:01:53Z
```

JSON reports contain the same statistics and incident information:

```json
{
  "statistics": {
    "total_valid": 5532,
    "total_malformed": 468,
    "level_counts": {
      "ERROR": 820,
      "FATAL": 271,
      "INFO": 3335,
      "WARN": 1106
    }
  },
  "incidents": []
}
```

The JSON example is shortened and does not show all services and incidents.

## Incident rules

GhostOps currently detects these incident types:

| Type | Detection |
| --- | --- |
| `FATAL_CRASH` | Any `FATAL` entry |
| `DATABASE_FAILURE` | `ERROR` or `FATAL` messages containing database, SQL, or connection pool terms |
| `PAYMENT_FAILURE` | `ERROR` or `FATAL` messages containing payment gateway or bad gateway terms |
| `APPLICATION_EXCEPTION` | `ERROR` or `FATAL` messages containing `exception` |
| `RESOURCE_PRESSURE` | `WARN` messages about high memory or a threshold |
| `RATE_LIMIT_WARNING` | `WARN` messages containing rate limit terms |
| `DEPRECATED_API` | `WARN` messages containing deprecated API terms |

A log entry can match more than one rule. For example, a fatal database error is included in both `FATAL_CRASH` and `DATABASE_FAILURE`.

Matches are grouped by incident type and service. Each group stores the number of matches, highest severity, earliest timestamp, and latest timestamp.

## Project structure

```text
cmd/                    command-line entry point
internal/parser/        line parsing and validation
internal/reader/        streaming file processing
internal/analyzer/      statistics and incident grouping
internal/incidentRule/  incident matching rules
internal/reporter/      text and JSON output
internal/domain/        shared models and severity values
```

## Design notes

- The file is processed with `bufio.Scanner` and is not loaded fully into memory.
- Only statistics and grouped incidents are kept in memory.
- Malformed lines are recoverable input problems. File and output errors stop execution.
- Services and incidents are sorted before text output so reports are deterministic.
- Incident rules use a small interface so more rules can be added without changing the analyzer.

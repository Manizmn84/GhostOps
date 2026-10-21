package enums

type Severity string

const (
	Info  Severity = "INFO"
	Warn  Severity = "WARN"
	Error Severity = "ERROR"
	Fatal Severity = "FATAL"
)

func IsSeverity(input string) bool {
	severity := Severity(input)

	switch severity {
	case Info, Warn, Error, Fatal:
		return true
	default:
		return false
	}
}

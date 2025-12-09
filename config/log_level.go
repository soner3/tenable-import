package config

import (
	"fmt"
	"strings"
)

// LogLevel definiert die möglichen Log-Level der Anwendung
type LogLevel int

// Mögliche Log-Level
const (
	Trace LogLevel = iota
	Debug
	Info
	Warn
	Error
)

// DefaultLogLevel ist das Standard-Log-Level
const DefaultLogLevel = Info

// String-Repräsentationen der Log-Level
var level = map[LogLevel]string{
	Trace: "TRACE",
	Debug: "DEBUG",
	Info:  "INFO",
	Warn:  "WARN",
	Error: "ERROR",
}

// String implementiert das Stringer-Interface
// und gibt die String-Repräsentation des Log-Levels zurück
func (l LogLevel) String() string {
	return level[l]
}

// ParseLogLevel parst einen String case-insensitive zu einem LogLevel
func ParseLogLevel(s string) (LogLevel, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "trace":
		return Trace, nil
	case "debug":
		return Debug, nil
	case "info":
		return Info, nil
	case "warn", "warning":
		return Warn, nil
	case "error":
		return Error, nil
	default:
		return Info, fmt.Errorf("ungültiges Log-Level: %q (erlaubt: trace, debug, info, warn, error)", s)
	}
}

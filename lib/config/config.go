package config

import (
	"io"
	"log"
)

// Version ist die aktuelle Version der Anwendung
const Version = "0.0.1"

// DefaultEnv ist die Standardumgebung, wenn keine Umgebungsvariable gesetzt ist
const DefaultEnv = "dev"

// disabledLogger ist ein Logger, der keine Ausgaben macht
var disabledLogger = log.New(io.Discard, "", 0)

// Config enthält die Konfiguration der Anwendung
// Es wird empfohlen, diese Konfiguration über Umgebungsvariablen zu setzen
type Config struct {
	Version     string
	Env         Environment
	LogLevel    LogLevel
	TraceLogger *log.Logger
	DebugLogger *log.Logger
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
}

// NewConfig erstellt eine neue Instanz von Config mit der Standardversion
func NewConfig() *Config {
	return &Config{
		Version: Version,
	}
}

// SetupLoggers konfiguriert die Logger basierend auf dem Log-Level
// Logger, die unter dem konfigurierten Level liegen, werden deaktiviert
func (cfg *Config) SetupLoggers(traceLogger, debugLogger, infoLogger, warnLogger, errorLogger *log.Logger) {
	switch cfg.LogLevel {
	case Trace:
		// Alle Logger aktiv
		cfg.TraceLogger = traceLogger
		cfg.DebugLogger = debugLogger
		cfg.InfoLogger = infoLogger
		cfg.WarnLogger = warnLogger
		cfg.ErrorLogger = errorLogger
	case Debug:
		// Debug und höher
		cfg.TraceLogger = disabledLogger
		cfg.DebugLogger = debugLogger
		cfg.InfoLogger = infoLogger
		cfg.WarnLogger = warnLogger
		cfg.ErrorLogger = errorLogger
	case Info:
		// Info und höher
		cfg.TraceLogger = disabledLogger
		cfg.DebugLogger = disabledLogger
		cfg.InfoLogger = infoLogger
		cfg.WarnLogger = warnLogger
		cfg.ErrorLogger = errorLogger
	case Warn:
		// Warn und höher
		cfg.TraceLogger = disabledLogger
		cfg.DebugLogger = disabledLogger
		cfg.InfoLogger = disabledLogger
		cfg.WarnLogger = warnLogger
		cfg.ErrorLogger = errorLogger
	case Error:
		// Nur Error
		cfg.TraceLogger = disabledLogger
		cfg.DebugLogger = disabledLogger
		cfg.InfoLogger = disabledLogger
		cfg.WarnLogger = disabledLogger
		cfg.ErrorLogger = errorLogger
	}
}

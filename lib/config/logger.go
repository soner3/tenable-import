package config

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
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
func (c *Config) ParseLogLevel(s string) (LogLevel, error) {
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
		return DefaultLogLevel, fmt.Errorf("ungültiges Log-Level: %q (erlaubt: trace, debug, info, warn, error)", s)
	}
}

func (c *Config) SetupLogger() {
	logLevel, err := c.ParseLogLevel(c.LogLevel.String())
	if err != nil {
		log.Fatalf("Fehler beim Parsen des Log-Levels: %v", err)
	}
	c.LogLevel = logLevel

	err = os.MkdirAll(c.LogFilePath, os.ModePerm)
	if err != nil {
		log.Fatalf("Log Verzeichnis konnte nicht erstellt werden: %v\n", err)
	}

	y, m, d := time.Now().Date()
	fName := fmt.Sprintf("%d%02d%02d%s.log", y, m, d, c.AppName)
	fmt.Println(filepath.Join(c.LogFilePath, fName))

	f, err := os.Create(filepath.Join(c.LogFilePath, fName))
	if err != nil {
		log.Fatalf("Log Datei konnte nicht erstellt werden: %v\n", err)
	}
	defer f.Close()

	stdWriter := io.MultiWriter(os.Stdout, f)
	logFlags := log.Ldate | log.Ltime | log.Lmicroseconds

	traceLogger := log.New(stdWriter, "TRACE\t", logFlags)
	debugLogger := log.New(stdWriter, "DEBUG\t", logFlags)
	infoLogger := log.New(stdWriter, "INFO\t", logFlags)
	warnLogger := log.New(stdWriter, "WARN\t", logFlags)
	errorLogger := log.New(stdWriter, "ERROR\t", logFlags)

	c.SetupLoggers(traceLogger, debugLogger, infoLogger, warnLogger, errorLogger)
}

// SetupLoggers konfiguriert die Logger basierend auf dem Log-Level
// Logger, die unter dem konfigurierten Level liegen, werden deaktiviert
func (c *Config) SetupLoggers(traceLogger, debugLogger, infoLogger, warnLogger, errorLogger *log.Logger) {
	switch c.LogLevel {
	case Trace:
		// Alle Logger aktiv
		c.TraceLogger = traceLogger
		c.DebugLogger = debugLogger
		c.InfoLogger = infoLogger
		c.WarnLogger = warnLogger
		c.ErrorLogger = errorLogger
	case Debug:
		// Debug und höher
		c.TraceLogger = disabledLogger
		c.DebugLogger = debugLogger
		c.InfoLogger = infoLogger
		c.WarnLogger = warnLogger
		c.ErrorLogger = errorLogger
	case Info:
		// Info und höher
		c.TraceLogger = disabledLogger
		c.DebugLogger = disabledLogger
		c.InfoLogger = infoLogger
		c.WarnLogger = warnLogger
		c.ErrorLogger = errorLogger
	case Warn:
		// Warn und höher
		c.TraceLogger = disabledLogger
		c.DebugLogger = disabledLogger
		c.InfoLogger = disabledLogger
		c.WarnLogger = warnLogger
		c.ErrorLogger = errorLogger
	case Error:
		// Nur Error
		c.TraceLogger = disabledLogger
		c.DebugLogger = disabledLogger
		c.InfoLogger = disabledLogger
		c.WarnLogger = disabledLogger
		c.ErrorLogger = errorLogger
	}
}

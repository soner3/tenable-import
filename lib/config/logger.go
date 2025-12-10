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
func (c *Config) ParseLogLevel(s string) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "trace":
		c.LogLevel = Trace
	case "debug":
		c.LogLevel = Debug
	case "info":
		c.LogLevel = Info
	case "warn", "warning":
		c.LogLevel = Warn
	case "error":
		c.LogLevel = Error
	default:
		c.LogLevel = DefaultLogLevel
	}
}

// InitLogger konfiguriert die Logger basierend auf dem Log-Level und erstellt die Log-Datei
// im angegebenen Verzeichnis. Diese Funktion sollte so früh wie möglich in der Anwendung
// aufgerufen werden.
//
// Die Funktion erstellt automatisch das Log-Verzeichnis falls es nicht existiert und
// generiert eine Log-Datei im Format YYYYMMDD-AppName-N.log, wobei N eine fortlaufende
// Nummer für mehrere Starts am selben Tag ist.
//
// Rückgabewert: Eine Cleanup-Funktion zum Freigeben von Ressourcen.
// Fehlerverhalten: Bei Dateisystem-Fehlern wird die Anwendung beendet.
// Log-Level: Bei ungültigen Werten wird DefaultLogLevel verwendet.
func (c *Config) InitLogger() func() {
	c.ParseLogLevel(c.LogLevel.String())

	err := os.MkdirAll(c.LogFilePath, os.ModePerm)
	if err != nil {
		log.Fatalf("Log Verzeichnis konnte nicht erstellt werden: %v\n", err)
	}

	y, m, d := time.Now().Date()
	pattern := filepath.Join(c.LogFilePath, fmt.Sprintf("%d%02d%02d-%s-*.log", y, m, d, c.AppName))
	matches, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatalf("Fehler beim Suchen alter Log-Dateien: %v\n", err)
	}

	run := 1

	if len(matches) > 0 {
		run = len(matches) + 1
	}

	fName := fmt.Sprintf("%d%02d%02d-%s-%d.log", y, m, d, c.AppName, run)

	f, err := os.Create(filepath.Join(c.LogFilePath, fName))
	if err != nil {
		log.Fatalf("Log Datei konnte nicht erstellt werden: %v\n", err)
	}

	stdWriter := io.MultiWriter(os.Stdout, f)
	logFlags := log.Ldate | log.Ltime | log.Lmicroseconds

	traceLogger := log.New(stdWriter, "TRACE\t", logFlags)
	debugLogger := log.New(stdWriter, "DEBUG\t", logFlags)
	infoLogger := log.New(stdWriter, "INFO\t", logFlags)
	warnLogger := log.New(stdWriter, "WARN\t", logFlags)
	errorLogger := log.New(stdWriter, "ERROR\t", logFlags)

	c.configLogger(traceLogger, debugLogger, infoLogger, warnLogger, errorLogger)
	return func() {
		defer f.Close()
	}
}

// SetupLoggers konfiguriert die Logger basierend auf dem Log-Level
// Logger, die unter dem konfigurierten Level liegen, werden deaktiviert
func (c *Config) configLogger(traceLogger, debugLogger, infoLogger, warnLogger, errorLogger *log.Logger) {
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

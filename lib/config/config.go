package config

import (
	"io"
	"log"
)

const (
	// Version der Anwendung
	Version = "1.0.0"
	// DefaultLogLevel ist das Standard-Log-Level
	DefaultLogLevel = Debug
	// DefaultEnv ist die Standardumgebung, wenn keine Umgebungsvariable gesetzt ist
	DefaultEnv = Dev
	// DefaultAppName ist der Standardname der Anwendung
	DefaultAppName = ""
	// DefaultLogFilePath ist der Standardpfad für Log-Dateien
	DefaultLogFilePath = "./logs/"
)

// disabledLogger ist ein Logger, der keine Ausgaben macht
var disabledLogger = log.New(io.Discard, "", 0)

// Config enthält die Konfiguration der Anwendung
// Es wird empfohlen, diese Konfiguration über Umgebungsvariablen zu setzen
type Config struct {
	Version     string
	Env         Environment
	AppName     string
	LogLevel    LogLevel
	LogFilePath string
	TraceLogger *log.Logger
	DebugLogger *log.Logger
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
}

// NewConfig erstellt eine neue Config mit Standardwerten
// Dabei werden Version, Umgebung und Log-Level gesetzt
// AppName wird leer gelassen und sollte vom Aufrufer gesetzt werden
func NewConfig() Config {
	return Config{
		Version:     Version,
		Env:         DefaultEnv,
		LogLevel:    DefaultLogLevel,
		AppName:     DefaultAppName,
		LogFilePath: DefaultLogFilePath,
	}
}

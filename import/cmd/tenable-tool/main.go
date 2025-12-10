package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/soner3/tenable-import/import/internal/config"
)

func main() {
	cfg := &config.AppConfig{}

	envFlag := flag.String("env", cfg.Env.String(), "Anwendungsumgebung (dev, qa, prod)")
	logLevelFlag := flag.String("loglevel", cfg.LogLevel.String(), "Log-Level (trace, debug, info, warn, error)")
	flag.String("dsn", cfg.DSN, "Datenbank-Verbindungszeichenfolge (Datasource Name - DSN)")
	flag.String("tenable-api-key", cfg.TenableAPIKey, "Tenable API-Schlüssel für die Authentifizierung")
	flag.Parse()

	env, err := cfg.ParseEnvironment(*envFlag)
	if err != nil {
		log.Fatalf("Fehler beim Parsen der Umgebung: %v", err)
	}
	cfg.Env = env

	logLevel, err := cfg.ParseLogLevel(*logLevelFlag)
	if err != nil {
		log.Fatalf("Fehler beim Parsen des Log-Levels: %v", err)
	}
	cfg.LogLevel = logLevel

	y, m, d := time.Now().Date()
	f, err := os.Create(fmt.Sprintf("%d%02d%02dTenableTool.log", y, m, d))
	if err != nil {
		log.Fatalf("Log Datei konnte nicht erstellt werden: %v\n", err)
		panic(err)
	}
	defer f.Close()

	stdWriter := io.MultiWriter(os.Stdout, f)
	errWriter := io.MultiWriter(os.Stderr, f)
	logFlags := log.Ldate | log.Ltime | log.Lmicroseconds

	traceLogger := log.New(stdWriter, "TRACE\t", logFlags)
	debugLogger := log.New(stdWriter, "DEBUG\t", logFlags)
	infoLogger := log.New(stdWriter, "INFO\t", logFlags)
	warnLogger := log.New(errWriter, "WARN\t", logFlags)
	errorLogger := log.New(errWriter, "ERROR\t", logFlags)

	cfg.SetupLoggers(traceLogger, debugLogger, infoLogger, warnLogger, errorLogger)

	cfg.InfoLogger.Printf("Anwendung gestartet in Umgebung: %q mit Log-Level: %q", cfg.Env, cfg.LogLevel)
	cfg.DebugLogger.Printf("Log-Datei erstellt: %q", f.Name())

}

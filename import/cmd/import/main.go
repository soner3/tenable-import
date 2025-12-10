package main

import (
	"flag"

	"github.com/soner3/tenable-import/import/internal/config"
)

func main() {
	app := config.NewAppConfig()

	env := ""
	logLevel := ""

	flag.StringVar(&env, "env", app.Env.String(), "Anwendungsumgebung (dev, qa, prod)")
	flag.StringVar(&logLevel, "log-level", app.LogLevel.String(), "Log-Level (trace, debug, info, warn, error)")
	flag.StringVar(&app.DSN, "dsn", app.DSN, "Datenbank-Verbindungszeichenfolge (Datasource Name - DSN)")
	flag.StringVar(&app.TenableAPIKey, "tenable-api-key", app.TenableAPIKey, "Tenable API-Schlüssel für die Authentifizierung")
	flag.StringVar(&app.AppName, "app-name", app.AppName, "Name der Anwendung")
	flag.Parse()

	app.SetupLogger()
	var err error
	app.Env, err = app.ParseEnvironment(env)

	if err != nil {
		app.ErrorLogger.Fatalf("Fehler beim Parsen der Umgebung: %v", err)
	}

	app.LogLevel, err = app.ParseLogLevel(logLevel)
	if err != nil {
		app.ErrorLogger.Fatalf("Fehler beim Parsen des Log-Levels: %v", err)
	}

	app.InfoLogger.Printf("Anwendung gestartet in Umgebung: %q mit Log-Level: %q", app.Env, app.LogLevel)
}

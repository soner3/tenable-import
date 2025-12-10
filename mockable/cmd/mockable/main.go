package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const Version = "1.0.0"
const Port = 8080
const ENV = "dev"

type config struct {
	Port        int
	ENV         string
	Version     string
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

func main() {
	cfg := &config{
		Port:    Port,
		ENV:     ENV,
		Version: Version,
	}

	flag.IntVar(&cfg.Port, "port", Port, "Port auf dem der Server läuft")
	flag.StringVar(&cfg.ENV, "env", ENV, "Environment (dev|prod)")
	flag.Parse()

	cfg.InfoLogger = log.New(log.Writer(), "INFO\t", log.Ldate|log.Ltime)
	cfg.ErrorLogger = log.New(log.Writer(), "ERROR\t", log.Ldate|log.Ltime)

	cfg.InfoLogger.Printf("Starte Server auf Port %d in %s Umgebung", cfg.Port, cfg.ENV)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), cfg.initRoutes()); err != nil {
		cfg.ErrorLogger.Fatalf("Fehler beim Starten des Servers: %v", err)
	}

}

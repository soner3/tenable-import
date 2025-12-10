package config

import "strings"

// Environment definiert die möglichen Anwendungsumgebungen
type Environment int

// Mögliche Umgebungen
const (
	Dev Environment = iota
	Qa
	Prod
)

// String-Repräsentationen der Umgebungen
var env = map[Environment]string{
	Dev:  "dev",
	Qa:   "qa",
	Prod: "prod",
}

// String implementiert das Stringer-Interface
// und gibt die String-Repräsentation der Umgebung zurück
func (e Environment) String() string {
	return env[e]
}

// ParseEnvironment parst einen String case-insensitive zu einer Environment
func (c *Config) ParseEnvironment(s string) {
	c.DebugLogger.Printf("Umgebung %q wird geparst", s)
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "dev", "development":
		c.Env = Dev
	case "qa", "test", "testing":
		c.Env = Qa
	case "prod", "production":
		c.Env = Prod
	default:
		c.WarnLogger.Printf("Ungültige Umgebung: %q, Standardumgebung %q wird verwendet", s, DefaultEnv)
		c.Env = DefaultEnv
	}
	c.InfoLogger.Printf("Umgebung gesetzt auf %q", c.Env)
}

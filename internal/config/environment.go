package config

import (
	"fmt"
	"strings"
)

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
func ParseEnvironment(s string) (Environment, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "dev", "development":
		return Dev, nil
	case "qa", "test", "testing":
		return Qa, nil
	case "prod", "production":
		return Prod, nil
	default:
		return Dev, fmt.Errorf("ungültige Umgebung: %q (erlaubt: dev, qa, prod)", s)
	}
}

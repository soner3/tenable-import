package config

import "github.com/soner3/tenable-import/lib/config"

type AppConfig struct {
	config.Config
	DSN           string
	TenableAPIKey string
}

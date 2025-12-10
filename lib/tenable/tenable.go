package tenable

import (
	"time"

	"github.com/soner3/tenable-import/lib/config"
)

const baseURL = "https://cloud.tenable.com"
const maxRetrys = 5
const defaultWaitSeconds = 5
const statusCheckInterval = 10 * time.Second

// TenableClient ist der Client für die Tenable API
type TenableClient struct {
	Config              *config.Config
	BaseURL             string
	MaxRetrys           int
	WaitSeconds         int
	StatusCheckInterval time.Duration
	ApiKey              string
}

// NewTenableClient erstellt eine neue Instanz von TenableClient
func NewTenableClient(cfg *config.Config) *TenableClient {
	return &TenableClient{
		Config:              cfg,
		BaseURL:             baseURL,
		MaxRetrys:           maxRetrys,
		WaitSeconds:         defaultWaitSeconds,
		StatusCheckInterval: statusCheckInterval,
	}
}

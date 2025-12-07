package tenable

// Tenable repräsentiert einen Tenable-Client mit API-Schlüssel und Basis-URL
type Tenable struct {
	APIKey  string
	BaseURL string
}

// BaseUrl ist die Standard-Basis-URL für die Tenable Cloud API
const BaseUrl = "https://cloud.tenable.com"

// NewTenableClient erstellt einen neuen Tenable-Client
// mit dem angegebenen API-Schlüssel und der Basis-URL
func NewTenableClient(apiKey, baseURL string) *Tenable {
	return &Tenable{
		APIKey:  apiKey,
		BaseURL: BaseUrl,
	}
}

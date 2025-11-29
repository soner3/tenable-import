package repository

// PiaRepository definiert die Methoden für die Interaktion mit PIA-Assets
type PiaRepository interface {
	GetAllAssets() string
}

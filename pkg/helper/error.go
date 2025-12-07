package helper

import (
	"fmt"
	"runtime/debug"
	"time"
)

// Error repräsentiert einen benutzerdefinierten Fehler mit zusätzlicher Kontextinformation
type Error struct {
	// Innerer Fehler, der eingewickelt wird
	Inner error
	// Aussagekräftige Fehlermeldung
	Message string
	// Stack-Trace zum Zeitpunkt der Fehlererstellung
	StackTrace string
	// Zeitpunkt der Fehlererstellung
	CreatedAt time.Time
	// Zusätzliche Kontextinformationen
	Ctx map[string]any
}

// Error implementiert das error-Interface
func WrapError(inner error, msgf string, args ...any) *Error {
	return &Error{
		Inner:      inner,
		Message:    fmt.Sprintf(msgf, args...),
		StackTrace: string(debug.Stack()),
		CreatedAt:  time.Now(),
		Ctx:        make(map[string]any),
	}
}

// Error implementiert das error-Interface
func (e *Error) Error() string {
	return e.Message
}

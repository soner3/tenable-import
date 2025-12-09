package db

import (
	"database/sql"

	_ "github.com/microsoft/go-mssqldb"
	"github.com/soner3/tenable-import/helper"
)

// CreateSQLServerConnection stellt eine Verbindung zu einer SQL Server-Datenbank her
func CreateSQLServerConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		err = helper.WrapError(err, "Fehler bei der Verbindung mit der DB:\n\t%q", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		err = helper.WrapError(err, "Fehler beim Pingen der DB:\n\t%q", err)
		return nil, err
	}
	return db, nil
}

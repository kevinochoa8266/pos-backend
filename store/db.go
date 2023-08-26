package store

import (
	_ "modernc.org/sqlite"
	"database/sql"
)

// Passes in the url, gets the connection, and checks if it is valid with a Ping.
func GetConnection(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbUrl); if err != nil {
		return nil, err
	}
	err = db.Ping(); if err != nil {
		return nil, err
	}
	return db, nil
}

func CloseConnection(db *sql.DB) (error) {
	err := db.Close(); if err != nil {
		return err
	}
	return nil
}


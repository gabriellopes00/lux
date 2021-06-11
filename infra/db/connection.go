package db

import (
	"database/sql"
)

func ConnectPg() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=psql dbname=helpy sslmode=disable")
	if err != nil {
		return nil, err
	} else if err = db.Ping(); err != nil {
		defer db.Close()
		return nil, err
	}

	return db, nil
}

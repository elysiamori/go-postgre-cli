package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DBConnect() (*sql.DB, error) {
	conn := ""

	db, err := sql.Open("postgres", conn)

	// best practice to close database

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	//fmt.Println("Database has been connected")

	return db, nil
}

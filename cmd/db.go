package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func connect() error {
	var err error

	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", cfg.PgHost, cfg.PgPort, cfg.PgBase, cfg.PgUser, cfg.PgPassword))
	if err != nil {
		return err
	}

	return nil
}

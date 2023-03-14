package main

import (
	"database/sql"
	"fiveWordsADay/log"
	_ "modernc.org/sqlite"
)

func openDB() {
	db, err := sql.Open("sqlite", "data.sqlite")
	if err != nil {
		log.LogError("Could not open database", err)
	}
	defer log.CloseAndLog(db, "Could not close database")
}

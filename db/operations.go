package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

// Database schema
const create_db_statement = `
	CREATE TABLE IF NOT EXISTS activity (
	  id TEXT NOT NULL PRIMARY KEY,
	  time TEXT NOT NULL,
	  distance REAL,
	  heartrate INTEGER,
	  power INTEGER
	  );`

func createDbSchema(db *sql.DB) {
	var err error

	_, err = db.Exec(create_db_statement)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, create_db_statement)
	} else {
		fmt.Println("SQLite database is ready.")
	}
}

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

const db_file string = "scratch.db"

func init_database() {
	var err error
	db, err := sql.Open("sqlite", db_file)

	if err != nil {
		log.Fatal(err)
	}

	create_db_statement := `
	CREATE TABLE IF NOT EXISTS scratch (
	  id INTEGER NOT NULL PRIMARY KEY,
	  time DATETIME NOT NULL,
	  distance float
	  );`

	_, err = db.Exec(create_db_statement)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, create_db_statement)
	} else {
		fmt.Println("SQLite database is ready.")
	}
}

func main() {

	init_database()
	fmt.Println("Hello World")
}

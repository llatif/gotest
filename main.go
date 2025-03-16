package main

import (
	"database/sql"
	"fmt"
	"os"
)

const db_file string = "scratch.db"               // SQLite database file name
var dbConnection, _ = sql.Open("sqlite", db_file) // SQLite database connection

func main() {

	// Check if the database file exists
	_, err := os.Stat(db_file)

	if os.IsExist(err) {
		fmt.Println("Database exists, proceeding to write data.")
	} else {
		db.createDbSchema(dbConnection) // Create database and schema
	}

	//singleActivity := newActivity("12:20", 23.4, 159, 200)
	//insertActivity(dbConnection, singleActivity)

	activity.fetchActivities(dbConnection)
}

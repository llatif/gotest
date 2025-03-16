package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

const db_file string = "scratch.db"               // SQLite database file name
var dbConnection, _ = sql.Open("sqlite", db_file) // SQLite database connection

type activity struct {
	id        string
	time      string
	distance  float32
	heartrate int
	power     int
}

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

func insertActivity(db *sql.DB, activity *activity) {
	// Insert activity into database

	insertActivityStatement :=
		`INSERT INTO activity (
			id, time, distance, heartrate, power) 
			VALUES (?, ?, ?, ?, ?);`

	statement, err := db.Prepare(insertActivityStatement)

	if err != nil {
		log.Fatalf("Error in activity format before entering into database: %q\n", err)
	}

	statement.Exec(activity.id, activity.time, activity.distance, activity.heartrate, activity.power)

	fmt.Printf("New activity stored with ID: %s\n, Time: %s\n, Distance: %f\n, Heartrate: %v\n, Power: %v\n",
		activity.id, activity.time, activity.distance, activity.heartrate, activity.power)

}

func fetchActivities(db *sql.DB) {
	// Return all activities

	row, err := db.Query("SELECT * from activity ORDER BY time")

	if err != nil {
		log.Fatalf("Error retrieving activity records: %q\n", err)
	}

	for row.Next() {
		var id string
		var time string
		var distance float32
		var heartRate int
		var power int

		row.Scan(&id, &time, &distance, &heartRate, &power)
		fmt.Printf("Activity ID: %s\n, Time: %s\n, Distance: %f\n, Heartrate: %v\n, Power: %v\n",
			id, time, distance, heartRate, power)
	}
}

func newActivity(time string, distance float32, heartrate int, power int) *activity {
	// Creates a new activity ready to be inserted into the database

	activityId := uuid.New()

	activityDetails := activity{
		id:        activityId.String(),
		time:      time,
		distance:  distance,
		heartrate: heartrate,
		power:     power}

	return &activityDetails
}

func main() {

	// Check if the database file exists
	_, err := os.Stat(db_file)

	if os.IsExist(err) {
		fmt.Println("Database exists, proceeding to write data.")
	} else {
		createDbSchema(dbConnection) // Create database and schema
	}

	//singleActivity := newActivity("12:20", 23.4, 159, 200)
	//insertActivity(dbConnection, singleActivity)

	fetchActivities(dbConnection)
}

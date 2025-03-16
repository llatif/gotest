package activity

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

type activity struct {
	id        string
	time      string
	distance  float32
	heartrate int
	power     int
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

package db

import (
	"database/sql"
	"fmt"
)

// init() : a function to setup the db by filling it with db schema needed or whatever
func Init() {
	db, err := sql.Open("sqlite", "./waypoints.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Always test the connection
	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("db connect success...")

	// creating new table
	query := `
		CREATE TABLE IF NOT EXISTS waypoints (
			waypoint_id INTEGER PRIMARY KEY AUTOINCREMENT,
			waypoint_name TEXT NOT NULL,
			waypoint_lat REAL NOT NULL,
			waypoint_lng REAL NOT NULL,
			waypoint_alt REAL NOT NULL
		);
		`

	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Created table or table existed...")
}

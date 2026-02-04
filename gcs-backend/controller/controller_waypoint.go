package controller

import (
	"database/sql"
	"gin-quickstart/models"

	_ "modernc.org/sqlite"
)

// get all
func GetAllWaypoints() []models.Waypoint {
	/**
	TO-DO
	1. connect to sqlite
	2. get all
	*/
	var waypoints []models.Waypoint

	db, err := sql.Open("sqlite", "./waypoints.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT waypoint_id, waypoint_name, waypoint_lat, waypoint_lng, waypoint_alt FROM waypoints`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var wp models.Waypoint
		err := rows.Scan(&wp.Waypoint_id, &wp.Waypoint_name, &wp.Waypoint_lat, &wp.Waypoint_lng, &wp.Waypoint_alt)
		if err != nil {
			panic(err)
		}
		waypoints = append(waypoints, wp)
	}

	return waypoints
}

// insert list of waypoints into db
func InsertWaypoints(waypoints []models.Waypoint) int {
	db, err := sql.Open("sqlite", "./waypoints.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// start transactions
	tx, err := db.Begin()

	// insert into the db
	stmt, err := tx.Prepare(`
		INSERT INTO waypoints (waypoint_name, waypoint_lat, waypoint_lng, waypoint_alt)
		VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// execute each
	for _, wp := range waypoints {
		_, err = stmt.Exec(wp.Waypoint_name, wp.Waypoint_lat, wp.Waypoint_lng, wp.Waypoint_alt)
		if err != nil {
			tx.Rollback()
			panic(err)
		}
	}

	// commit transaction
	tx.Commit()

	return len(waypoints)
}

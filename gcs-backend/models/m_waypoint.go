package models

// bikin strukutr dari waypoint
type Waypoint struct {
	waypoint_id   string
	waypoint_name string
	waypoint_lat  float64
	waypoint_lng  float64
	waypoint_alt  int
}

// get all
func getAllWaypoints() {
	/**
	TO-DO
	1. connect to sqlite
	2. insert
	*/
}

// insert list of waypoints into db
func insertWaypoints(waypoints []Waypoint) {

}

// put
func (w Waypoint) putWaypoint() {

}

// delete
func (w Waypoint) deleteWaypoint() {

}

package models

type Waypoint struct {
	Waypoint_id   int     `json:"waypoint_id"`
	Waypoint_name string  `json:"waypoint_name"`
	Waypoint_lat  float64 `json:"waypoint_lat"`
	Waypoint_lng  float64 `json:"waypoint_lng"`
	Waypoint_alt  float32 `json:"waypoint_alt"`
}

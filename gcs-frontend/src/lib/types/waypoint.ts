// this writable is using pub-sub architecture to fill and get the waypoints from 
// the map and back into the sidebar
import { writable } from "svelte/store"

export interface Waypoint {
    waypoint_id : number,
    waypoint_name : string,
    waypoint_lat : number,
    waypoint_lng :  number,
    waypoint_alt : number
}

// export interface InsertWaypointRequestDTO {
//     waypoint_name : string,
//     waypoint_lat : number,
//     waypoint_lng :  number,
//     waypoint_alt : number
// }

export const waypoints = writable<Waypoint[]>([]);
export const payloadWaypoints = writable<Waypoint[]>([]); // this is the array for the send ones
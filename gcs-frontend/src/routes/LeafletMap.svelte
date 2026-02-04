<svelte:head>
    <link rel="stylesheet" href=
"https://unpkg.com/leaflet@1.6.0/dist/leaflet.css"
    integrity="sha512-xwE/Az9zrjBIphAcBb3F6JVqxf46+CDLwfLMHloNu6KEQCAWi6HcDUbeOfBIptF7tcCzusKFjFw2yuvEpDL9wQ=="
    crossorigin="" />
</svelte:head>

<!-- script for fetching the  the leaflet JavaScript file -->
<script lang="ts" src="https://unpkg.com/leaflet@1.6.0/dist/leaflet.js"
    integrity="sha512-gZwIG9x3wUXg2hdXF6+rVkLF/0Vi9U8D2Ntg4Ga5I5BZpVkVxlJWbSQtXPSiUTtC0TjtGOmxa1AJPuV0CPthew=="
    crossorigin="">

    import AksantaraMarker from '$lib/images/aksantara.svg';
    import { waypoints } from '$lib/types/waypoint';
    import type { Waypoint } from '$lib/types/waypoint';
    import Sidebar from '$lib/components/Sidebar.svelte';
    import { onMount } from "svelte";
	import type { Map } from 'leaflet';

    // function for get the altitude using fetch api to another website
    async function fetchWaypoints() {
        // Construct the API URL
        const apiUrl = `http://127.0.0.1:9090/waypoints`;
        console.log(`Fetching waypoints : ${apiUrl}`);

        const response = await fetch(apiUrl);

        if (!response.ok) {
            console.log(`Response not ok : ${response.status}`)
        }

        const body = await response.json();
        console.log("Fetch waypoints success. Waypoints:", body);
        return body;
    }

    // ========================= init (get the db first) ==========================
    onMount(async () => {
        // init (get the db at first)
        const data = await fetchWaypoints();
        waypoints.set(data);
    })
    
    // ======================== map functionalities =============================
    let map : Map;
    
    onMount(async () => {

        // function for get the altitude using fetch api to another website
        async function getAltitude(lat : number, lng : number) {
            // Construct the API URL
            const apiUrl = `https://api.open-meteo.com/v1/elevation?latitude=${lat}&longitude=${lng}`;
            console.log(apiUrl)

            const response = await fetch(apiUrl);

            if (!response.ok) {
                console.log(`Response not ok : ${response.status}`)
            }

            const value = await response.json()
            console.log(`Fetch altitude success. Altitude : ${value.elevation}`);
            return value.elevation
        }

        
        // Initialize the map
        const L = await import("leaflet");
        map = L.map('map')
        
        const aksantaraIcon = L.icon({
            iconUrl: AksantaraMarker,
            iconSize: [50, 50],        
            iconAnchor: [16, 32],     
            popupAnchor: [0, -32]
        });

        // Get the tile layer from OpenStreetMaps
        L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        
            // Specify the maximum zoom of the map
            maxZoom: 19,
            
            // Set the attribution for OpenStreetMaps
            attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        }).addTo(map);
        
        // Set the view of the map
        // with the latitude, longitude and the zoom value
        // map.setView([48.8584, 2.2945], 16);
        
        // Set the map view to the user's location
        // Uncomment below to set map according to user location
        map.locate({setView: true, maxZoom: 16});
        
        // Show a market at the position of the Eiffel Tower
        let eiffelMarker = L.marker([48.8584, 2.2945]).addTo(map);
        
        // Bind popup to the marker with a popup
        eiffelMarker.bindPopup("Eiffel Tower").openPopup();

        // set click function to map
        map.on('click', onMapClick);

        // define functions for clicking
        async function onMapClick(e : any) {
            
            // split the data
            console.log("You clicked the map at " + e.latlng);
            const lat = e.latlng.lat;
            const long = e.latlng.lng;
            const alt = await getAltitude(lat, long)

            // create new waypoint
            const new_wp : Waypoint = {
                waypoint_id : "",
                waypoint_name : "Dummy Test Waypoint",
                waypoint_lat : lat,
                waypoint_lng : long,
                waypoint_alt : alt
            }

            // add into waypoints
            waypoints.update(wps => [...wps, new_wp]);

            // add marker
            L.marker([lat, long], { icon: aksantaraIcon })
                .addTo(map)
                .bindPopup(new_wp.waypoint_name);
        }

    })


</script>

<div class="map-wrapper">
    <div id="map" >
    </div>
    <div id="sidebar-container">
        <Sidebar />
    </div>
</div>

<!-- css lah pokoknya -->
<style>
    #map {
        width: 100%;
        height: 100%;
    }
    
    .map-wrapper {
        position: relative;
        height: 100vh;
        width: 100vw;
    }
</style>
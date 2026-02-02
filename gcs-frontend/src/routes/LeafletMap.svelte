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

    import { onMount } from "svelte";

    let map;

    
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
            console.log(`Altitude : ${value.elevation}`);
            alert("Altitude Fetch Success");
        }


        // define functions for clicking
        async function onMapClick(e : any) {
            console.log("You clicked the map at " + e.latlng);
            const lat = e.latlng.lat;
            const long = e.latlng.lng;
            await getAltitude(lat, long)
        }

        // Initialize the map
        const L = await import("leaflet");
        map = L.map('map')
        
        // Get the tile layer from OpenStreetMaps
        L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        
            // Specify the maximum zoom of the map
            maxZoom: 19,
            
            // Set the attribution for OpenStreetMaps
            attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        }).addTo(map);
        
        // Set the view of the map
        // with the latitude, longitude and the zoom value
        map.setView([48.8584, 2.2945], 16);
        
        // Set the map view to the user's location
        // Uncomment below to set map according to user location
        map.locate({setView: true, maxZoom: 16});
        
        // Show a market at the position of the Eiffel Tower
        let eiffelMarker = L.marker([48.8584, 2.2945]).addTo(map);
        
        // Bind popup to the marker with a popup
        eiffelMarker.bindPopup("Eiffel Tower").openPopup();

        // set click function to map
        map.on('click', onMapClick);
        
    })


</script>

<div id="map" style="width: 960px; height: 60vh"></div>

<!-- css lah pokoknya -->
<style>
    /* css stuff */
</style>
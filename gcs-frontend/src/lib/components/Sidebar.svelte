
<script lang="ts">
    import { payloadWaypoints, waypoints } from "$lib/types/waypoint";
    import type { Waypoint } from "$lib/types/waypoint";

    // function to submit to backend
    async function submitWaypoints() {
        // construct the API URL
        const apiUrl = `http://127.0.0.1:9090/waypoints`;
        console.log(`Submitting waypoints : ${apiUrl}`);

        // constructing the body
        const content : Waypoint[] = $payloadWaypoints;
        console.log("Body : ", content);

        const response = await fetch(apiUrl, {
            method: "POST",
            body : JSON.stringify(content)
        });

        if (!response.ok) {
            console.log(`Response not ok : ${response.status}`)
        }

        const body = await response.json();
        console.log(`Submit waypoints success. Waypoints : ${body}`);
    }

    async function EditWaypointName(lat: number, lng : number, alt : number, new_name:string) {
        let target_wp = $waypoints.find(wp => (wp.waypoint_lat === lat) && (wp.waypoint_lng === lng) && (wp.waypoint_alt === alt));
        if (target_wp) {
            let old_name = target_wp.waypoint_name;
            // edit waypoints
            waypoints.update(wps =>
                wps.map(wp =>
                    (wp.waypoint_lat === lat) && (wp.waypoint_lng === lng) && (wp.waypoint_alt === alt) ? { ...wp, waypoint_name: new_name } : wp
                )
            );
            console.log(`Update waypoints success. Changed name from ${old_name} to ${name}`);
            // edit payloadWaypoints
            payloadWaypoints.update(wps =>
                wps.map(wp =>
                    (wp.waypoint_lat === lat) && (wp.waypoint_lng === lng) && (wp.waypoint_alt === alt) ? { ...wp, waypoint_name: new_name } : wp
                )
            );
            console.log(`Update payload waypoints success. Changed name from ${old_name} to ${name}`);
        }
        console.log(`Not find any such waypoint in the current store. Lat : ${lat}. Long : ${lng}. Alt : ${alt}.`);
    }

    // function to delete from the current waypoint
    async function deleteWaypoint(lat: number, lng : number, alt : number) {
        waypoints.update(wps =>
            wps.filter(wp => (wp.waypoint_lat !== lat) && (wp.waypoint_lng !== lng) && (wp.waypoint_alt !== alt))
        );
        payloadWaypoints.update(wps =>
            wps.filter(wp => (wp.waypoint_lat !== lat) && (wp.waypoint_lng !== lng) && (wp.waypoint_alt !== alt))
        );
        console.log(`Deleted waypoint. Lat : ${lat}. Long : ${lng}. Alt : ${alt} `);
    };

</script>

<div class="sidebar-container">

    <!-- simple title -->
    <h1 class="waypoint-title">Waypoints List</h1>

    <!-- save -->
    <div class="save-container" on:click={() => submitWaypoints()}>
        Save
    </div>

    <!-- below are the waypoints list -->
    <div class="waypoint-block">
        {#if $waypoints.length === 0}
            <div class="no-waypoints-container">
                No waypoints in db
            </div>
        {:else}
            <div class="waypoint">
                {#each $waypoints as wp}

                <!-- the controls -->
                <div class="waypoint-detail">
                    <!-- the waypoints -->
                    <div class="waypoint-data">
                        <br/>
                        <div><b>Waypoint ID : </b>{wp.waypoint_id}</div>
                        <div><b>Name : </b>
                            <input value={wp.waypoint_name} on:input={(e) => EditWaypointName(wp.waypoint_lat, wp.waypoint_lng, wp.waypoint_alt,( e.target as HTMLInputElement).value)}/>
                        </div>
                        <div><b>Latitude (degree): </b> {wp.waypoint_lat}</div>
                        <div><b>Longitude (degree) : </b>{wp.waypoint_lng}</div>
                        <div><b>Altitude (m) : </b>{wp.waypoint_alt}</div>
                    </div>

                    <!-- save change -->
                    <div class="waypoint-save">
                        Save
                    </div>

                    <!-- delete -->
                    <div class="waypoint-delete" on:click={() => deleteWaypoint(wp.waypoint_lat, wp.waypoint_lng, wp.waypoint_alt)}>
                        Delete
                    </div>
                </div>


                {/each}
    
            </div>
        {/if}
    </div>
    
</div>

<style>
    .sidebar-container {
        position: absolute;
        padding: 1%;
        top: 1.5rem;
        left: 3rem;
        z-index: 1000;
        justify-content: center;
        align-items: center;
        background-color: white;
        border-radius: 1pc;
        max-height: 80vh; 
        width: 35%;   

        display:  flex;
        flex-direction: column;
    }

    .save-container {
        display: flex;
        height: 100%;
        width: 100%;
        justify-content: center;
        align-items: center;
        background-color: green;
        font-weight: 650;
        padding: 2%;
        border-radius: 1rem;
    }

    .waypoint-detail {
        display: flex;
        justify-content: space-between;
        align-items: center;
        height: fit-content;
    }

    .waypoint-delete {
        padding: 2%;
        background-color: red;
        display: flex;
        height: 100%;
        width: 80%;
        justify-content: center;
        align-items: center;
        border: 2px solid black;
        border-radius: 1em;
    }

    .waypoint-save {
        padding: 2%;
        background-color: green;
        display: flex;
        height: 100%;
        width: 80%;
        justify-content: center;
        align-items: center;
        border: 2px solid black;
        border-radius: 1em;
    }

    .waypoint-data {
        margin: 2%;
    }

    .waypoint-block {
        scroll-behavior: auto;
        width: 100%;
        height: 100%;
        overflow-y: auto;     
        flex: 1; 
    }

    .waypoint-title {
        font-weight: 700;
    }

    .no-waypoints-container {
        margin: 2%;
        font-weight: 600;
        display: flex;
        justify-content: center;
        align-items: center;
    }

</style>
# Medium Ground Control Station 
A ground control station for doing connection, controlling, and monitoring of Unmanned Aerial Vehicles (UAVs) and/or Remote Control Vehicles (RCVs). This is made in the purpose of answering one of the questions in THT2 from Ca-RSC26 Aksantara ITB.

# Functionalities
1. Mission Planning & Control: Allows operators to define, upload, and update waypoints and flight paths, as well as take manual control of the UAV. <b>(DONE)</b>
2. Real-time Telemetry & Monitoring: Displays live data on the vehicle’s position, altitude, speed, and vital statistics. <b>(PLANNED)</b>
3. Payload Management: Manages and controls onboard sensors, cameras, and other equipment.  <b>(PLANNED)</b>
4. Communication Link: Maintains a secure, two-way data link (radio or satellite) between the operator and the unmanned vehicle.  <b>(PLANNED)</b>
5. Video Feed & Processing: Receives and displays live video streams, often with recording capabilities. <b>(PLANNED)</b>
6. System Configuration & Maintenance: Facilitates setting up autopilot parameters, updating firmware, and conducting pre-flight checks. <b>(PLANNED)</b>
7. Data Recording & Analysis: Records flight data for post-flight analysis, debriefing, and replay <b>(PLANNED)</b>

# Tech Stack
- Node JS (v22.14.0)
- Bun (v1.3.5)
- Leaflet.js
- Docker Engine + Docker (intinya docker)

# How to run
## Frontend
1. `bun install`
2. `bun run dev --open` (flag --open digunakan untuk menampilkan log yang lebih rinci)
## Backend
2. `go run main.go`

# How to run using Docker
1. `docker compose up --build` atau `docker compose up` kalau mau lebih cepet.

# Author
> Brian A. Hadian (13523048)

# Sources
- https://daily.dev/blog/svelte-for-beginners-a-guide
- https://lionontheweb.medium.com/sveltekit-and-bun-development-environment-with-docker-f04eac5e7adb
- https://www.mapbox.com/mapbox-gljs
- https://docs.mapbox.com/mapbox-gl-js/example/simple-map/
- https://www.geeksforgeeks.org/javascript/using-leaflet-js-to-show-maps-in-a-webpage/
https://stackoverflow.com/questions/62374265/svelte-with-leaflet
- https://leafletjs.com/examples/quick-start/
- https://leafletjs.com/reference.html
- https://stackoverflow.com/questions/13142635/how-can-i-create-an-object-based-on-an-interface-file-definition-in-typescript
- https://stackoverflow.com/questions/67873142/how-to-listen-the-state-changes-in-svelte-like-useeffect

# Notes

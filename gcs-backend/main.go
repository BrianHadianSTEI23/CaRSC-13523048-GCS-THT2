package main

import (
	"fmt"
	"gin-quickstart/controller"
	"gin-quickstart/db"
	"gin-quickstart/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func main() {
	db.Init()

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/waypoints", func(c *gin.Context) {

		var waypoints []models.Waypoint = controller.GetAllWaypoints()

		fmt.Println(waypoints)

		c.JSON(200, waypoints)
	})

	router.POST("/waypoints", func(c *gin.Context) {
		var waypoints []models.Waypoint

		// Parse JSON body
		if err := c.ShouldBindJSON(&waypoints); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Insert into DB
		count := controller.InsertWaypoints(waypoints)
		fmt.Println(count)

		c.JSON(201, gin.H{
			"waypoint_count": count,
		})
	})

	router.Run(":9090")
}

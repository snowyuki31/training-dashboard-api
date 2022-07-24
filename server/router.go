package server

import (
	"snowyuki31/training-dashboard-api/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	v1 := router.Group("v1")
	{
		activityGroup := v1.Group("activity")
		{
			activity := new(controllers.ActivityController)
			activityGroup.GET("/latest", activity.RetrieveLatest)
		}
	}

	return router
}

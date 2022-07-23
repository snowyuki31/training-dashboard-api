package controllers

import (
	"net/http"
	"snowyuki31/training-dashboard-api/service"

	"github.com/gin-gonic/gin"
)

type ActivityController struct{}

func (a ActivityController) RetrieveLatest(c *gin.Context) {

	activityService := service.ActivityService{}

	activity := activityService.LoadData()

	c.JSON(http.StatusOK, activity)
}

package controllers

import (
	"net/http"

	"snowyuki31/training-dashboard-api/service"

	"github.com/gin-gonic/gin"
)

type StatisticsController struct{}

func (s StatisticsController) Retrieve(c *gin.Context) {

	statisticsService := service.StatisticsService{}

	statistics := statisticsService.GetStatistics()

	c.JSON(http.StatusOK, statistics)

}

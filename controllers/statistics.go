package controllers

import (
	"net/http"

	"snowyuki31/training-dashboard-api/model"

	"github.com/gin-gonic/gin"
)

type StatisticsController struct{}

func (s StatisticsController) Retrieve(c *gin.Context) {

	statistics := model.GetStatistics()

	c.JSON(http.StatusOK, statistics)

}

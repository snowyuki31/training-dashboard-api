package controllers

import (
	"io/ioutil"
	"net/http"
	"snowyuki31/training-dashboard-api/model"

	"github.com/gin-gonic/gin"
)

type ActivityController struct{}

func (a ActivityController) Retrieve(c *gin.Context) {
	id := c.Param("id")

	activity := model.LoadData(id)
	mean := activity.CalcMean()
	max := activity.CalcMax()

	c.JSON(http.StatusOK, gin.H{"activity": activity.Trackpoint, "mean": mean, "max": max})
}

func (a ActivityController) ActivityList(c *gin.Context) {
	files, err := ioutil.ReadDir("data")
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		paths = append(paths, file.Name()[9:len(file.Name())-4])
	}

	c.JSON(http.StatusOK, gin.H{"activities": paths})
}

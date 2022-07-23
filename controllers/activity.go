package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ActivityController struct{}

func (a ActivityController) RetrieveLatest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Latest Activity"})
}

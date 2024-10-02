package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func HomeHandler(c *gin.Context) {
    query_param := c.Query("param")
	data := gin.H{
		"message": query_param,
		"count":   21,
	}
	c.JSON(http.StatusOK, data)
	
}
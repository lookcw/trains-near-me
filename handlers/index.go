package handlers

import (
    "github.com/gin-gonic/gin"
	"net/http"

)


func HomeHandler(c *gin.Context) {
    query_param := c.Query("param")
	data := gin.H{
		"message": query_param,
		"count":   42,
	}
	c.JSON(http.StatusOK, data)
	
}
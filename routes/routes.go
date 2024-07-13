package routes

import (
    "github.com/gin-gonic/gin"
    "trains-near-me/handlers"
)


func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/",handlers.HomeHandler)
	return router
	
}


package routes

import (
	"trains-near-me/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/", handlers.HomeHandler)
	router.GET("/test", handlers.TestHandler)
	return router

}

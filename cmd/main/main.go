package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // Initialize Gin router
    router := gin.Default()

    // Define routes
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Welcome to Gin Framework!",
        })
    })

    // Start the server
    router.Run(":8080")
}
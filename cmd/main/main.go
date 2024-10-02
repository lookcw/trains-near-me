package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
	gtfs_rt "github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
	gtfs "github.com/artonge/go-gtfs"
	// "golang.org/x/exp/slices"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net/http"
	"runtime"
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
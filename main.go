package main

import (
	"trains-near-me/routes"
	)

func main() {
    // Initialize Gin router
    router := routes.SetupRouter()
    // Start the server
    router.Run(":8080")
}
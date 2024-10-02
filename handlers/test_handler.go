package handlers

import (
	"fmt"
	"log"
	"net/http"
	models "trains-near-me/models"

	gtfs "github.com/artonge/go-gtfs"
	"github.com/gin-gonic/gin"
)

func TestHandler(c *gin.Context) {

	g_feed, err := gtfs.Load("supplemental_mta_gtfs", nil)

	if err != nil {
		log.Fatal(err)
	}
	stopTimeMap := models.NewTripStopTimeMap(g_feed.StopsTimes)
	relevantTrips := models.GetTripIdsWithStop(stopTimeMap, "G20")
	for _, tripId := range relevantTrips {
		fmt.Printf("%s, ", tripId)
	}

	c.String(http.StatusOK, "Hello World!")

}

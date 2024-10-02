package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"

	gtfs_rt "github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
	"google.golang.org/protobuf/proto"
)

func getRelevantStopTimeUpdates(tripIds []string) []*gtfs_rt.TripUpdate_StopTimeUpdate {

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api-endpoint.mta.info/Dataservice/mtagtfsfeeds/nyct%2Fgtfs", nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	feed := gtfs_rt.FeedMessage{}
	err = proto.Unmarshal(body, &feed)
	if err != nil {
		log.Fatal(err)
	}

	for _, entity := range feed.Entity {
		tripUpdate := entity.GetTripUpdate()
		if len(tripUpdate.GetStopTimeUpdate()) > 0 {
			// if tripUpdate.GetStopTimeUpdate()[0] != nil && slices.Contains(stationsOfInterest, tripUpdate.GetStopTimeUpdate()[0].GetStopId()) {

			if tripUpdate.GetStopTimeUpdate()[0] != nil {
				// stopId:=tripUpdate.GetStopTimeUpdate()[0].GetStopId()
				// firstCharacter:=stopId[0:1]
				fmt.Printf("Stop ID: %s\n", tripUpdate.GetStopTimeUpdate()[0].GetStopId())
				fmt.Printf("Trip ID: %s\n", tripUpdate.GetTrip().GetTripId())
				fmt.Printf("Route ID: %s\n", tripUpdate.GetTrip().GetRouteId())
				fmt.Printf("Timestamp: %d\n", tripUpdate.GetStopTimeUpdate()[0].GetArrival().GetTime())
				fmt.Printf("Timestamp: %d\n", *tripUpdate.GetStopTimeUpdate()[0].StopSequence)
			}
			// if tripUpdate != nil && tripUpdate.StopTimeUpdate != nil && len(tripUpdate.StopTimeUpdate) > 0 && tripUpdate.StopTimeUpdate[0] != nil {
			// 	fmt.Printf("Stop ID: %s\n", tripUpdate.StopTimeUpdate[0].GetStopId())
			// }
		}
	}
}

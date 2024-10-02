package models

import (
	utils "trains-near-me/utils"

	"github.com/artonge/go-gtfs"
)

type TripStopTimeMap map[string][]gtfs.StopTime

func NewTripStopTimeMap(stopTimes []gtfs.StopTime) TripStopTimeMap {
	validLines := map[string]struct{}{
		"1": {}, "2": {}, "3": {}, "4": {},
		"5": {}, "6": {}, "7": {}, "A": {},
		"B": {}, "C": {}, "D": {}, "E": {},
		"F": {}, "G": {}, "J": {}, "L": {},
		"M": {}, "N": {}, "Q": {}, "R": {},
		"S": {}, "W": {},
	}
	tripIdToStopTimes := make(map[string][]gtfs.StopTime)

	for _, stopTime := range stopTimes {
		routeId := utils.GetRouteIDFromTripID(stopTime.TripID)
		_, isValidRouteId := validLines[routeId]
		if tripArray, exists := tripIdToStopTimes[stopTime.TripID]; !exists && isValidRouteId {
			tripIdToStopTimes[stopTime.TripID] = []gtfs.StopTime{stopTime}
		} else if isValidRouteId {
			tripIdToStopTimes[stopTime.TripID] = append(tripArray, stopTime)
		}
	}

	return tripIdToStopTimes
}

func GetTripIdsWithStop(tripStopTimeMap TripStopTimeMap, stopId string) []string {
	var tripsWithStop []string
	for tripId, stopTimeArray := range tripStopTimeMap {
		for _, stopTime := range stopTimeArray {
			if stopTime.StopID[:3] == stopId[:3] {
				tripsWithStop = append(tripsWithStop, tripId)
			}
		}
	}
	return tripsWithStop
}

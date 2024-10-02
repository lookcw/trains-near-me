package utils

import (
	"strings"
)

func GetRouteIDFromTripID(tripId string) string {
	firstPart := strings.Split(tripId, "..")[0]
	return strings.Split(firstPart, "_")[2]
}

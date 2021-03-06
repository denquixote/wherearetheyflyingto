// Package destinationfinder provides functionality for
// determining where a flight with a given callsign is
// destined.
package destinationfinder

import (
	"os"
)

/**
 * Retrieves the lat long of the destination (as a simple string, we're not interested in doing
 * any real processing with this, just using it as an index.
 **/
type DestinationFinder interface {
	GetDestinationFromCallsign(callsign string) (lat_long string, err error)
}

func GetDestinationFinder() DestinationFinder {
	findername := os.Getenv("WATFT_FINDER")
	switch findername {
	case "planefinder":
		return PlaneFinderDestinationFinder{}
	case "holidayextras":
		return HolidayExtrasDestinationFinder{}
	default:
		return PlaneFinderDestinationFinder{}
	}
	return nil
}

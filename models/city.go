package models

import (
	"math"
	"os"
	s "strings"
	"time"

	"github.com/gocarina/gocsv"
)

// City representation
type City struct {
	Name string `csv:"city_ascii"`
	Coordinate
	CountryCode string `csv:"iso2"`
	NotUsed     string `csv:"-"`
}

// Coordinate representation
type Coordinate struct {
	Lat float64 `csv:"lat"`
	Lng float64 `csv:"lng"`
}

var citiesDB = make(map[string]Coordinate)

// Kps (kilometers per second) is the MAX LEVEL speed a Boeing 777-300ER Engine
// can reach at a height of 35,000 ft
const Kps = 0.2583333333 // kilometer per second assuming max speed of (930 km/h)

// GetRealDistance calculates the distance between two points (given the
// latitude/longitude of those points). It is being used to calculate
// the distance between two locations using GeoDataSource (TM) prodducts
// Definitions:
// South latitudes are negative, east longitudes are positive
// Passed to function:
// lat1, lon1 = Latitude and Longitude of point 1 (in decimal degrees)
// lat2, lon2 = Latitude and Longitude of point 2 (in decimal degrees)
// unit = the unit you desire for results
// where: 'M' is statute miles (default)
// 'K' is kilometers
// 'N' is nautical miles
func GetRealDistance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}
	return dist
}

// GetTime calculates the difference of Final time and Initial time
func GetTime(t1, t2 time.Time) float64 {
	diff := t2.Sub(t1)
	return diff.Seconds()
}

// Math for calculating if transaction is pbysically (commenrcially, not military) possible
// Assuming the maximum distance somone can travel (in KM) k/s * s
func getPotentialDistance(kps float64, deltaSeconds float64) float64 {
	return (kps * deltaSeconds)
}

// real distance should always be = or <= to potential
// distance (ideal distance that you can travel with ideal conditions)
func isPosible(realDistance, potentialDistance float64) bool {
	if realDistance > potentialDistance {
		return false
	}
	return true
}

func init() {
	citiesFile, err := os.OpenFile("worldcities.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer citiesFile.Close()

	cities := []*City{}

	if err := gocsv.UnmarshalFile(citiesFile, &cities); err != nil { // Load cities from file
		panic(err)
	}

	for _, city := range cities {
		citiesDB[s.ToLower(city.Name+city.CountryCode)] = Coordinate{Lat: city.Lat, Lng: city.Lng}
	}
}

// GetCitiesDB returns a map with a key [CityName + CountryCode]
// and a value of type Coordinate
func GetCitiesDB() map[string]Coordinate {
	return citiesDB
}

package controllers
import (
  // "fmt"
  "math"
  "time"
  // s "strings"

  // u "github.com/shekodn/anti-fraud-api/utils"
  // "github.com/shekodn/anti-fraud-api/models"
)

// Boeing 777-300ER Engine
// Max Level Speed (at altitude)575 mph (930 km/h) at 35,000 ft (10,675 m), Mach 0.87
const Kps = 0.2583333333 // kilometer per second assuming max speed of (930 km/h)


// This routine calculates the distance between two points (given the
// latitude/longitude of those points)
// lat1, lon1 = Latitude and Longitude of point 1 (in decimal degrees)
// lat2, lon2 = Latitude and Longitude of point 2 (in decimal degrees)

// unit = the unit you desire for results:
// 'M' is statute miles (default)
// 'K' is kilometers
// 'N' is nautical miles

func getRealDistance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {

  const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1) * math.Sin(radlat2) + math.Cos(radlat1) * math.Cos(radlat2) * math.Cos(radtheta)

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

func getDeltaTime(t1, t2 time.Time) (float64) {
  diff := t2.Sub(t1)
  return diff.Seconds()
}

// Math for calculating if transaction is pbysically (commenrcially, not military) possible
// Assuming the maximum distance somone can travel (in KM) k/s * s
func getPotentialDistance(kps float64, deltaSeconds float64) float64{
  return(kps*deltaSeconds)
}

// This functions gets both:
// Potential Distance: Max distance someone can travel in ideal conditions
// Real Distance: Distance someone traveled between current TX and previous TX
func IsLegit(userId uint, cityName, countryCode, txTime string) bool {

  // // Cache logic
  // previousCity := s.ToLower(lastTransactions[previousTx.UserId].CityName)
  // previousCountryCode := s.ToLower(lastTransactions[previousTx.UserId].CountryCode)
  // previousKey := previousCity + previousCountryCode
  // previousTime, err := time.Parse(time.RFC822, lastTransactions[previousTx.UserId].Time)
  //
  // if err != nil {
  //   fmt.Println("Error:", err)
  //   return u.Message(false, "Unable to parse previousTime"), false
  // }
  //
  // currentCity := s.ToLower(cityName)
  // currentCountryCode := s.ToLower(countryCode)
  // currentKey :=  currentCity + currentCountryCode
  // currentTime, err := time.Parse(time.RFC822, txTime.Time)
  //
  // if err != nil {
  //   fmt.Println("Error:", err)
  //   return u.Message(false, "Unable to parse currentTime"), false
  // }
  //
  // previousCoordinate, isPresentPrevious := models.GetCitiesDB()[previousKey]
  // currentCoordinate, isPresentCurrent := models.GetCitiesDB()[currentKey]
  //
  // if !isPresentPrevious || !isPresentCurrent {
  //   return u.Message(false, "One of the cities doesn't exist"), false
  // }
  //
  // realDistance := getRealDistance(previousCoordinate.Lat, previousCoordinate.Lng, currentCoordinate.Lat, currentCoordinate.Lng, "K")
  // deltaTime := getDeltaTime(previousTime, currentTime) // seconds
  // potentialDistance := getPotentialDistance(Kps, deltaTime)
  //
  // fmt.Println("PREVIOUS TX: ", lastTransactions[currentTx.UserId])
  //
  // if isPosible(realDistance, potentialDistance) {
  //
  //   addLastTx(currentTx)
  //
  //   fmt.Println("NEW TX: ", lastTransactions[currentTx.UserId])
  //
  //   return u.Message(true, "Transaction is Possible"), true
  // }

  return true

}


// real distance should always be = or <= to potential
// distance (ideal distance that you can travel with ideal conditions)
func isPosible(realDistance, potentialDistance float64) bool {
  if (realDistance > potentialDistance) {
    return false
  }
  return true
}

// this will be the cache
// var lastTransactions = make(map[uint]Transaction)

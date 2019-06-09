package controllers

import (
  "fmt"
  "math"
  "time"
  "strconv"
  "encoding/json"
  s "strings"
  u "github.com/shekodn/anti-fraud-api/utils"
  "github.com/shekodn/anti-fraud-api/models"
)


type TxHelper struct {
  Time string
  Lat string
  Lng string
}


// Boeing 777-300ER Engine
// Max Level Speed (at altitude)575 mph (930 km/h) at 35,000 ft (10,675 m), Mach 0.87
const maxSpeed = 0.2583333333 // kilometer per second assuming max speed of (930 km/h)

// This routine calculates the distance between two points (given the
// latitude/longitude of those points)
// lat1, lon1 = Latitude and Longitude of point 1 (in decimal degrees)
// lat2, lon2 = Latitude and Longitude of point 2 (in decimal degrees)
// unit = the unit you desire for results: 'M' is statute miles (default), 'K' is kilometers, 'N' is nautical miles
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

// Math for calculating if transaction is pbysically (commercially, not military) possible
// Assuming the maximum distance somone can travel (in KM) k/s * s
func getPotentialDistance(kps float64, deltaSeconds float64) float64{
  return(kps*deltaSeconds)
}

// Aux func
func strCoordinatePairToCoordinate(strLat, strLng string) models.Coordinate {

  lat, err := strconv.ParseFloat(strLat, 64)

  if err != nil {
    panic(err)
  }

  lng, err := strconv.ParseFloat(strLng, 64)

  if err != nil {
    panic(err)
  }

  return models.Coordinate{Lat: lat, Lng: lng}
}


// real distance should always be = or <= to potential
// distance (distance that you can travel with ideal conditions)
func isPosible(realDistance, potentialDistance float64) bool {
  if (realDistance > potentialDistance) {
    return false
  }
  return true
}


func strToTime(strTime string) (time.Time, error) {

  currentTime, err := time.Parse(time.RFC822, strTime)

  return currentTime, err
}

// This functions gets both:
// Potential Distance: Max distance someone can travel in ideal conditions
// Real Distance: Distance someone traveled between current TX and previous TX
func IsLegit(userId uint, cityName, countryCode, txTime string) (map[string]interface{}, bool) {

  // If there is a previous key, value pair, compare it with the current info
  // (city, country and time) in order to determine if the transaction is
  // physically possible
  if strPreviousCoordinatePairAndTime, ok := models.GetPreviousCoordinate(userId); ok {

    auxPreviousTx := TxHelper{}
    json.Unmarshal([]byte(strPreviousCoordinatePairAndTime), &auxPreviousTx)

    // TODO: Refactor this
    previousCoordinate := strCoordinatePairToCoordinate(auxPreviousTx.Lat, auxPreviousTx.Lng)
    strPreviousTime := fmt.Sprint(auxPreviousTx.Time)
    previousTime, err := time.Parse(time.RFC822, strPreviousTime)

    if err != nil {
      fmt.Println("Error:", err)
      return u.Message(false, "Unable to parse previousTime"), false
    }

    currentKey :=  s.ToLower(cityName + countryCode)
    currentCoordinate, isCurrentCityinDB := models.GetCitiesDB()[currentKey]
    currentTime, err := time.Parse(time.RFC822, txTime)

    if err != nil {
      fmt.Println("Error:", err)
      return u.Message(false, "Unable to parse currentTime"), false
    }

    if !isCurrentCityinDB {
      return u.Message(false, "Current city doesn't exist"), false
    }

    realDistance := getRealDistance(previousCoordinate.Lat, previousCoordinate.Lng, currentCoordinate.Lat, currentCoordinate.Lng, "K")
    deltaTime := getDeltaTime(previousTime, currentTime) // seconds
    potentialDistance := getPotentialDistance(maxSpeed, deltaTime)

    if isPosible(realDistance, potentialDistance) {

      strCurrentLat := fmt.Sprint(currentCoordinate.Lat)
      strCurrentLng := fmt.Sprint(currentCoordinate.Lng)
      auxCurrentTx := &TxHelper{Lat: strCurrentLat, Lng: strCurrentLng, Time: txTime}

      json, err := json.Marshal(auxCurrentTx)
      if err != nil {
          fmt.Println(err)
          return u.Message(false, "Problem while doing the Unmarshal"), false
      }

      models.SetLastCoordinate(userId, string(json))

      return u.Message(true, "Transaction is Possible"), true

    } else {

      return u.Message(false, "Transaction is Fraudalent"), false
    }

  } else {

    // Since the is not a previous TX, we add the current transaction values to
    // the previous transactions DB (cache)
    return addCurrentKeyValuePairToCache(userId, cityName, countryCode, txTime)
  }
}


func addCurrentKeyValuePairToCache(userId uint, cityName, countryCode, txTime string) (map[string]interface{}, bool) {
  // CURRENT coordinate will become PREVIOUS coordinate
  previousKey := s.ToLower(cityName + countryCode)
  previousCoordinate, isPreviousCityinDB := models.GetCitiesDB()[previousKey]

  if !isPreviousCityinDB {
    return u.Message(false, "Current city/country doesn't exist"), false
  }

  strPreviousLat := fmt.Sprint(previousCoordinate.Lat)
  strPreviousLng := fmt.Sprint(previousCoordinate.Lng)

  auxTx := &TxHelper{Lat: strPreviousLat, Lng: strPreviousLng, Time: txTime}

  json, err := json.Marshal(auxTx)
  if err != nil {
      fmt.Println(err)
      return u.Message(false, "Problem while doing the Unmarshal"), false
  }

  models.SetLastCoordinate(userId, string(json))

  return u.Message(true, "Transaction is possible"), true
}

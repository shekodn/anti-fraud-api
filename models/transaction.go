package models

import (
  "time"
  "fmt"
  s "strings"
  u "github.com/shekodn/anti-fraud-api/utils"
  "github.com/jinzhu/gorm"
)

type Transaction struct {
  gorm.Model
  CityName string `json:"city_name"`
  CountryCode string `json:"country_code"` // country code according to the ISO2
  Time string `json:"time"`
  UserId uint `json:"user_id"`
}

var lastTransactions = make(map[uint]Transaction)


func getPotentialDistance(kps float64, deltaSeconds float64) float64{
  return(kps*deltaSeconds)
}

func isPosible(realDistance, potentialDistance float64) bool {
  if (realDistance > potentialDistance) {
    return false
  }
  return true
}

func isFraud(tx *Transaction) (map[string] interface{}, bool) {

  if _, ok := lastTransactions[tx.UserId]; !ok {
    lastTransactions[tx.UserId] = Transaction{
      CityName: tx.CityName,
      CountryCode: tx.CountryCode,
      UserId: tx.UserId,
      Time: tx.Time,
    }

    return u.Message(true, "success"), true
  }

  // Boeing 777-300ER Engine
  // https://www.google.com/search?q=velocity+boeing+777&oq=velocity+boeing+777&aqs=chrome..69i57j0l2.4438j0j4&sourceid=chrome&ie=UTF-8
  // Max Level Speed (at altitude)	575 mph (930 km/h) at 35,000 ft (10,675 m), Mach 0.87
  const Kps = 0.2583333333 // kilometer per second assuming max speed of (930 km/h)

  previousCity := s.ToLower(lastTransactions[tx.UserId].CityName)
  previousCountryCode := s.ToLower(lastTransactions[tx.UserId].CountryCode)
  previousKey := previousCity + previousCountryCode
  // previousTime, err := time.Parse(time.RFC3339, "2012-05-08T14:30:05.000Z")
  previousTime, err := time.Parse(time.RFC822, lastTransactions[tx.UserId].Time)

  fmt.Println("User's last valid tx was in: ", previousKey)

  if err != nil {
    fmt.Println(err)
    return u.Message(false, "Unable to parse previousTime"), false
  }

  currentCity := s.ToLower(tx.CityName)
  currentCountryCode := s.ToLower(tx.CountryCode)
  currentKey :=  currentCity + currentCountryCode
  // currentTime, err := time.Parse(time.RFC3339, "2012-05-08T15:03:05.000Z")
  currentTime, err := time.Parse(time.RFC822, tx.Time)

  if err != nil {
    fmt.Println(err)
    return u.Message(false, "Unable to parse currentTime"), false
  }

  previousCoordinate, isPresentPrevious := GetCitiesDB()[previousKey]
  currentCoordinate, isPresentCurrent := GetCitiesDB()[currentKey]

  if !isPresentPrevious || !isPresentCurrent {
    return u.Message(false, "One of the cities doesn't exist"), false
  }

  realDistance := GetRealDistance(previousCoordinate.Lat, previousCoordinate.Lng, currentCoordinate.Lat, currentCoordinate.Lng, "K")
  deltaTime := GetTime(previousTime, currentTime) // seconds
  potentialDistance := getPotentialDistance(Kps, deltaTime)

  // fmt.Printf("Previous: %s Current: %s \n", previousKey, currentKey)
  // fmt.Printf("Delta time (S): %f\n", deltaTime)
  // fmt.Printf("Real Distance in KM: %f\n", realDistance)
  // fmt.Printf("Potential Distance in KM (Best Case Scenario): %f\n", potentialDistance)
  // fmt.Printf("Is Possible?: %t\n", isPosible(realDistance, potentialDistance))


  if isPosible(realDistance, potentialDistance) {
    lastTransactions[tx.UserId] = Transaction{
      CityName: tx.CityName,
      CountryCode: tx.CountryCode,
      UserId: tx.UserId,
      Time: tx.Time,
    }
    fmt.Println("User's last valid tx is NOW: ", lastTransactions[tx.UserId])
    return u.Message(true, "Transaction is Possible"), true
  }

  return u.Message(false, "Transaction is Fraudalent"), false
}


//The FRAUD MS should be here
func (tx *Transaction) Validate() (map[string] interface{}, bool) {

    // https://golang.org/pkg/time/
    // RFC822:
    _, err := time.Parse(time.RFC822, tx.Time)

    if err != nil {
      fmt.Println("Err:", err)
      resp := u.Message(false, "Unable to parse time")
      return resp, false
    }

    if resp, ok := isFraud(tx); !ok {
      fmt.Println("tx.isFraud")
      return resp, false
    }

    return u.Message(true, "success"), true
}

func (tx *Transaction) Create() (map[string] interface {}) {

  if resp, ok := tx.Validate(); !ok {
    fmt.Println("TX couldn't be validated")
    return resp
  }

  GetDB().Create(tx)

  resp := u.Message(true, "success")
  resp["tx"] = tx
  return resp

}

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


// this will be the cache
var lastTransactions = make(map[uint]Transaction)

// Cache functions
// Checks if previous transaction exists in cache memory
func previousTxExist(currentTx *Transaction, m map[uint]Transaction) (bool) {
  _, ok := m[currentTx.UserId]
  return ok
}

// gets previous transaction
func getPreviousTx(currentTx *Transaction) (Transaction) {
  return lastTransactions[currentTx.UserId]
}

// Add current Transaction as last transaction to Cache
func addLastTx(currentTx *Transaction){
  lastTransactions[currentTx.UserId] = Transaction{
    CityName: currentTx.CityName,
    CountryCode: currentTx.CountryCode,
    UserId: currentTx.UserId,
    Time: currentTx.Time,
  }
}

func IsLegit(currentTx *Transaction, previousTx Transaction) (map[string] interface{}, bool) {

  previousCity := s.ToLower(lastTransactions[previousTx.UserId].CityName)
  previousCountryCode := s.ToLower(lastTransactions[previousTx.UserId].CountryCode)
  previousKey := previousCity + previousCountryCode
  previousTime, err := time.Parse(time.RFC822, lastTransactions[previousTx.UserId].Time)

  if err != nil {
    fmt.Println("Error:", err)
    return u.Message(false, "Unable to parse previousTime"), false
  }

  currentCity := s.ToLower(currentTx.CityName)
  currentCountryCode := s.ToLower(currentTx.CountryCode)
  currentKey :=  currentCity + currentCountryCode
  currentTime, err := time.Parse(time.RFC822, currentTx.Time)

  if err != nil {
    fmt.Println("Error:", err)
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

  fmt.Println("PREVIOUS TX: ", lastTransactions[currentTx.UserId])

  if isPosible(realDistance, potentialDistance) {
    lastTransactions[currentTx.UserId] = Transaction{
      CityName: currentTx.CityName,
      CountryCode: currentTx.CountryCode,
      UserId: currentTx.UserId,
      Time: currentTx.Time,
    }

    fmt.Println("NEW TX: ", lastTransactions[currentTx.UserId])

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

    if previousTxExist(tx, lastTransactions) {

      previousTx := getPreviousTx(tx)

      if resp, ok := IsLegit(tx, previousTx); !ok {
        fmt.Println("IS FRAUD!")
        return resp, false
      } else {
        // TX is not fraud
        return u.Message(true, "success"), true
      }
    } else {
      // add last tx because it doesn't exist in cache
      addLastTx(tx)
      // Returns TRUE because since TX was not in cache, you cannot flag it
      // as fraudulent
      return u.Message(true, "success"), true
    }
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

package models

import(
  "fmt"
  "github.com/go-redis/redis"
  // s "strings"

)

var client *redis.Client

func init() {
  //creates cache for key value pair for each TX
  client = redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
  })

  pong, err := client.Ping().Result()
  fmt.Println(pong, err)
}

func GetCache() *redis.Client {
  return client
}

/* Cache functions */

// Checks if previous coordinate exists in cache memory
func previousCoordinateExist(userId uint, client *redis.Client) (bool) {

  strUserId := fmt.Sprint(userId)

  _, err := client.Get(strUserId).Result()

  if err == redis.Nil {
    return false
  } else if err != nil {
    panic(err)
    // TODO: return false should I add a return here? No, right?
  } else {
    return true
  }
}

// Get previous coordinate
func getPreviousCoordinate(userId uint, client *redis.Client) string {

  strUserId := fmt.Sprint(userId)

  coordinate, err := client.Get(strUserId).Result()
  if err != nil {
    panic(err)
  }

  fmt.Println("User ID:", strUserId, "Coord:", coordinate)

  return coordinate
}

// Add current Transaction as last transaction to Cache
func setLastCoordinate(userId uint, coordinate Coordinate){

  strUserId := fmt.Sprint(userId)

  strCurrentCoordinates := fmt.Sprint(coordinate.Lat + coordinate.Lng)

  err := GetCache().Set(strUserId, strCurrentCoordinates, 0).Err()
  if err != nil {
	  panic(err)
  }
}




// if previousTxExist(tx, lastTransactions) {
//
//   previousTx := getPreviousTx(tx)
//
//   if resp, ok := IsLegit(tx, previousTx); !ok {
//     fmt.Println("IS FRAUD!")
//     return resp, false
//   } else {
//     // TX is not fraud
//     return u.Message(true, "success"), true
//   }
// } else {
//   // add last tx because it doesn't exist in cache
//   addLastTx(tx)
//   // Returns TRUE because since TX was not in cache, you cannot flag it
//   // as fraudulent
//   return u.Message(true, "success"), true
// }

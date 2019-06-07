package models

import(
  "fmt"
  "github.com/go-redis/redis"
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

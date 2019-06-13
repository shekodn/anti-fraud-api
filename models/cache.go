package models

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

var client *redis.Client

func init() {

	e := godotenv.Load()

	if e != nil {
		log.Println("ENV: ", e)
	}

	dbHost := os.Getenv("DB_HOST")
	cachePort := os.Getenv("CACHE_PORT")
	// cachePassword := os.Getenv("CACHE_PASSWORD")

	//Build connection string
	cacheUri := fmt.Sprintf("host=%s port=%s", dbHost, cachePort)
	log.Println(cacheUri)

	//creates cache for key value pair for each TX
	client = redis.NewClient(&redis.Options{
		Addr:     dbHost + ":" + cachePort,
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
func GetPreviousCoordinate(userId uint) (string, bool) {

	strUserId := fmt.Sprint(userId)

	coordinate, err := client.Get(strUserId).Result()

	if err == redis.Nil {
		return "", false
	} else if err != nil {
		panic(err)
		// TODO: return false should I add a return here? No, right?
	} else {
		return coordinate, true
	}
}

// Add current Transaction as last transaction to Cache
func SetLastCoordinate(userId uint, coordinatePair string) {

	strUserId := fmt.Sprint(userId)

	err := GetCache().Set(strUserId, coordinatePair, 0).Err()
	if err != nil {
		panic(err)
	}
}

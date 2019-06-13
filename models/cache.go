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

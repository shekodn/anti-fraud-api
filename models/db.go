package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()

	if e != nil {
		log.Println("ENV: ", e)
	}

	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("DB_HOST")

	//Build connection string
	// TODO: Password should not be logged
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	log.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)

	if err != nil {
		log.Print("Error:", err, "\n")
	}

	db = conn

}

// getDB returns gorm db
func getDB() *gorm.DB {
	return db
}

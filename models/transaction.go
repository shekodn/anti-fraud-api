package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/shekodn/anti-fraud-api/utils"
	s "strings"
	"time"
)

type Transaction struct {
	gorm.Model
	CityName    string `json:"city_name"`
	CountryCode string `json:"country_code"` // country code according to the ISO2
	Time        string `json:"time"`
	UserId      uint   `json:"user_id"`
}

func (tx *Transaction) Validate() (map[string]interface{}, bool) {

	// https://golang.org/pkg/time/
	// RFC822:
	_, err := time.Parse(time.RFC822, tx.Time)

	if err != nil {
		fmt.Println("Err:", err)
		resp := u.Message(false, "Unable to parse time")
		return resp, false
	}

	return u.Message(true, "success: TX was validated"), true

}

func (tx *Transaction) Create() map[string]interface{} {

	if resp, ok := tx.Validate(); !ok {
		fmt.Println("TX couldn't be validated")
		return resp
	}

	// Lowercase name and country before saving them in DB
	tx.CityName = s.ToLower(tx.CityName)
	tx.CountryCode = s.ToLower(tx.CountryCode)

	getDB().Create(tx)

	resp := u.Message(true, "success")
	resp["tx"] = tx
	return resp

}

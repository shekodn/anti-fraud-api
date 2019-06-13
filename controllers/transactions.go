package controllers

import (
	// "fmt"
	// "strconv"
	"encoding/json"
	"github.com/shekodn/anti-fraud-api/models"
	u "github.com/shekodn/anti-fraud-api/utils"
	"net/http"
)

var CreateTransaction = func(w http.ResponseWriter, r *http.Request) {
	//Grabs the id of the user that send the request
	tx := &models.Transaction{}

	err := json.NewDecoder(r.Body).Decode(tx)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	if result, ok := IsLegit(tx.UserId, tx.CityName, tx.CountryCode, tx.Time); !ok {
		w.WriteHeader(http.StatusUnavailableForLegalReasons)
		u.Respond(w, result)
		return
	}

	resp := tx.Create()

	w.WriteHeader(http.StatusCreated)
	u.Respond(w, resp)
}

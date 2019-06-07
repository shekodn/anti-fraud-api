package controllers

import(
  // "fmt"
  // "strconv"
  "net/http"
  "encoding/json"
  "github.com/shekodn/anti-fraud-api/models"
  u "github.com/shekodn/anti-fraud-api/utils"

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

    if !IsLegit(tx.UserId, tx.CityName, tx.CountryCode, tx.Time) {
      w.WriteHeader(http.StatusUnavailableForLegalReasons)
      u.Respond(w, u.Message(false, "Transaction is Fraudalent"))
      return
    }

    resp := tx.Create()

    w.WriteHeader(http.StatusCreated)
    u.Respond(w, resp)
}

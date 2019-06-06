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

    resp := tx.Create()

    // isSuccess, err := strconv.ParseBool(resp["status"])
    //
    // if err != nil {
    //   fmt.Println("Error: ", err)
    //   return
    // }
    //
    // if isSuccess {
    //   w.WriteHeader(http.StatusCreated)
    // } else{
    //   w.WriteHeader(http.StatusBadRequest)
    // }
    w.WriteHeader(http.StatusCreated)
    u.Respond(w, resp)
}

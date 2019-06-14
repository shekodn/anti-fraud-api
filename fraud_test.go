// Reference: https://blog.alexellis.io/golang-writing-unit-tests/
package main

// import (
//   "testing"
//   "bytes"
//   "net/http"
//   "net/http/httptest"
//   "encoding/json"
//   // "fmt"
//   "github.com/shekodn/anti-fraud-api/models"
//   "github.com/shekodn/anti-fraud-api/controllers"
//
// )
//
// func TestCreateTransaction(t *testing.T) {
//
// 	var jsonStr = []byte(`{ "user_id" : 2, "city_name" : "Berlin", "country_code" : "de", "time" : "05 Aug 12 02:47 UTC"}`)
//
// 	req, err := http.NewRequest("POST", "/new", bytes.NewBuffer(jsonStr))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(controllers.CreateTransaction)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusCreated {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusCreated)
// 	}
// }

// func TestIsFraud1(t *testing.T) {
//   var tx1 models.Transaction
//   var tx2 models.Transaction
//   var tx3 models.Transaction
//
//   var jsonBerlin =  []byte(`{"user_id" : 2, "city_name" : "Berlin", "country_code" : "de", "time" : "05 Aug 12 02:47 UTC"}`)
//   var jsonToronto = []byte(`{"user_id" : 2, "city_name" : "toronto", "country_code" : "ca", "time" : "05 Aug 12 14:30 UTC"}`)
//   var jsonCancun =  []byte(`{"user_id" : 2, "city_name" : "cancun", "country_code" : "mx", "time" : "05 Aug 12 15:03 UTC"}`)
//
//   err := json.Unmarshal(jsonBerlin, &tx1)
//   if err != nil {
//     panic(err)
//   }
//
//   err = json.Unmarshal(jsonToronto, &tx2)
//   if err != nil {
//     panic(err)
//   }
//
//   err = json.Unmarshal(jsonCancun, &tx3)
//   if err != nil {
//     panic(err)
//   }
//
//   _, isLegit1 := models.IsLegit(&tx2, tx1)
//
//   if isLegit1 != true {
//     t.Errorf("Transaction is supposed to be legit: got %v want %v",
//       isLegit1, true)
//   }
//
//   _, isLegit2 := models.IsLegit(&tx3, tx2)
//
//   if isLegit2 != false {
//     t.Errorf("Transaction is supposed to NOT be legit: got %v want %v",
//       isLegit2, false)
//   }
// }

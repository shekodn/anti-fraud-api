// Reference: https://blog.alexellis.io/golang-writing-unit-tests/
package main

import (
  "testing"
  "bytes"
  "net/http"
  "net/http/httptest"
  // "fmt"
  // "github.com/shekodn/anti-fraud-api/models"
  "github.com/shekodn/anti-fraud-api/controllers"

)

func TestCreateTransaction(t *testing.T) {

	var jsonStr = []byte(`{ "user_id" : 2, "city_name" : "Berlin", "country_code" : "de", "time" : "05 Aug 12 02:47 UTC"}`)

	req, err := http.NewRequest("POST", "/new", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.CreateTransaction)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	// expected := `{ "user_id" : 2, "city_name" : "Berlin", "country_code" : "de", "time" : "05 Aug 12 02:47 UTC"}`
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }
}


//
// func TestIsFraud1(t *testing.T) {
//
//
// }
//
// func TestIsFraud2(t *testing.T) {
//
//
// }

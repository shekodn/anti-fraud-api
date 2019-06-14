package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shekodn/anti-fraud-api/controllers"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Run server: go build -o app && ./app
// Try requests: curl http://127.0.0.1:8000/
func main() {
	log.Info("Initialize service...")

	r := mux.NewRouter()
	// Application controller
	r.HandleFunc("/healthz", controllers.GetHealthz).Methods("GET")
	r.HandleFunc("/readyz", controllers.GetReadyz).Methods("GET")

	// Transaction controller
	r.HandleFunc("/new", controllers.CreateTransaction).Methods("POST")

	log.Info("Service is ready to listen and serve.")

	err := http.ListenAndServe(":8000", r)

	if err != nil {
		log.Println("Error:", err)
	}
}

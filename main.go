package main

import(
  "fmt"
  "log"
  "os"
  "net/http"
  "github.com/shekodn/anti-fraud-api/controllers"
  "github.com/gorilla/mux"
)


func main () {
  r := mux.NewRouter()
  r.HandleFunc("/new", controllers.CreateTransaction).Methods("POST")
  r.HandleFunc("/__heartbeat__", controllers.GetHeartbeat).Methods("GET")

  port := os.Getenv("PORT")

  if port == "" {
    port = "8000"
  }

  fmt.Println("Listening in " + port)
  err := http.ListenAndServe(":" + port, r)

  if err != nil {
    log.Println("Error:", err)
  }
}

package controllers

import (
	"net/http"
)

var GetHeartbeat = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I'm alive"))
}

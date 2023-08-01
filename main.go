package main

import (
	//"encoding/json"
	//"io"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type crate struct {
  Name string `json:"name"`
}


func getName(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-type", "application/json")
  test := crate {
    Name: "test",
  }
  json.NewEncoder(w).Encode(test)
}
func main() {
  router := mux.NewRouter()
  router.HandleFunc("/{crate_name}", getName).Methods("GET")
  http.ListenAndServe(":8000", router)
}

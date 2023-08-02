package main

import (
	//"encoding/json"
	//"io"
	"encoding/json"
	"net/http"
  "log"

	"github.com/gorilla/mux"
)

type crate struct {
  Name string `json:"name"`
}

func getName(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  name := vars["crate_name"]
  w.Header().Set("Content-type", "application/json")
  json.NewEncoder(w).Encode(name)
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/crate/{crate_name}", getName).Methods("GET")
  log.Fatal(http.ListenAndServe(":8080", router))
}

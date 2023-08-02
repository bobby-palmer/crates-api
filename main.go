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

func getInfo(name string) crate {
  return crate {
    Name: name,
  }
}

func getName(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  name := vars["name"]
  w.Header().Set("Content-type", "application/json")
  json.NewEncoder(w).Encode(getInfo(name))
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/crate/{name}", getName).Methods("GET")
  log.Fatal(http.ListenAndServe(":8080", router))
}

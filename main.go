package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type crate struct {
  Name string `json:"name"`
  Downloads string `json:"downloads"`
}

func extract(key string, body *http.Response) string {
  return "test"
}

func getInfo(name string) crate {
  url := "https://crates.io/api/v1/crates/" + name
  res, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  log.Print(res)
  downloads := extract("downloads", res)

  return crate {
    Name: name,
    Downloads: downloads,
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

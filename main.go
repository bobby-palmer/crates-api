package main

import (
	"encoding/json"
	"log"
	"net/http"
  "io/ioutil"
  "regexp"

	"github.com/gorilla/mux"
)

type endpoint struct {
  Version int `json:"schemaVersion"`
  Label string `json:"label"`
  Message string `json:"message"`
  Color string `json:"color"`
}

func extract(key string, body string) string {
  prefix := `"` + key + `":"?`
  pattern := regexp.MustCompile(prefix + `(\w+)`)
  matches := pattern.FindSubmatch([]byte(body))
  if matches == nil {
    return ""
  }
  return string(matches[1])
}

func getInfo(name string) endpoint {
  url := "https://crates.io/api/v1/crates/" + name
  res, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    log.Fatal(err)
  }
  downloads := extract("downloads", string(body))
  return endpoint {
    Version: 1,
    Label: "Downloads",
    Message: downloads,
    Color: "orange",
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

package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	. "github.com/tbxark/g4vercel"
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

func Handler(w http.ResponseWriter, r *http.Request) {
  server := New()

  server.GET("/downloads/:name", func(context *Context) {
    context.JSON(400, getInfo(context.Param("name")))
  })

  server.GET("/", func(context *Context) {
    context.JSON(400, H {
      "message":"SUCCESS!",
    })
  })

  server.Handle(w, r)
}

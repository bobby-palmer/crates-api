package handler

import (
	"io/ioutil"
	"log"
	"net/http"
  "regexp"

	. "github.com/tbxark/g4vercel"
)

func separate(num string) string {
  length := len(num)
  if length <= 3 {
    return num
  }
  return separate(num[0:length - 3]) + "," + num[length - 3:]
}


func extract(key string, body string) string {
  prefix := `"` + key + `":"?`
  pattern := regexp.MustCompile(prefix + `(\w+)`)
  matches := pattern.FindSubmatch([]byte(body))
  if matches == nil {
    return ""
  }
  return separate(string(matches[1]))
}

func getInfo(name string) H {
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
  return H {
    "schemaVersion": 1,
    "label": "Downloads",
    "message": downloads,
    "color": "orange",
  }
}

func serveBadge(name string) *http.Response {
  url := "https://img.shields.io/endpoint?url=https%3A%2F%2Fcrates-api.vercel.app%2Fdownloads%2F" + name
  res, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  return res
}

func Handler(w http.ResponseWriter, r *http.Request) {
  server := New()

  server.GET("/downloads/:name", func(context *Context) {
    context.JSON(200, getInfo(context.Param("name")))
  })

  server.GET("/", func(context *Context) {
    context.JSON(400, H {
      "message":"SUCCESS!",
    })
  })

  server.GET("/crates/badge/:name", func(context *Context) {
    context.JSON(200, serveBadge(context.Param("name")))
  })

  server.Handle(w, r)
}

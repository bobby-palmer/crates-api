package main

import(
  "io"
  "net/http"
)

func main() {
  handler := func(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "testing!")
  }
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8000", nil)
}

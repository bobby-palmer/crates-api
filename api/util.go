package handler

import (
  "regexp"
)

func extract(key string, body string) string {
  prefix := `"` + key + `":"?`
  pattern := regexp.MustCompile(prefix + `(\w+)`)
  matches := pattern.FindSubmatch([]byte(body))
  if matches == nil {
    return ""
  }
  return separate(string(matches[1]))
}

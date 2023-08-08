package handler

import (
	"testing"
)


func TestSeparate(t *testing.T) {
  result := separate("1000000")
  expected := "1,000,000"
  if result != expected {
    t.Errorf("got %q, wanted %q", result, expected)
  }
}

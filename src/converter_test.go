package main

import "testing"

func TestConvertToExcadecimal(t *testing.T) {

  result := ConvertToExadecimal(1)

  if result != "1" {
    t.Errorf("Conversion of 1 is WRONG!: expected %s, got %s", "1", result)
  } else {
    t.Logf("Conversion of 1 is 1: OK")
  }
}

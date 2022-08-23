package validators

import (
  "reflect"
)

func CheckArrayForType(arr []interface{}, expectedType string) bool {
  for _, item := range arr {
    itemType := reflect.TypeOf(item).String()
    if itemType != expectedType {
      return false
    }
  }
  return true
}

func SliceContains(arr []string, matchString string) bool {
  for _, a := range arr {
    if a == matchString {
      return true
    }
  }
  return false
}

package validators

import (
  "reflect"
  "time"
)

const DATE_FORMAT = "2006-01-02"
const DATETIME_FORMAT = time.RFC3339

func CheckArrayForType(arr []interface{}, expectedType string) bool {
  for _, item := range arr {
    itemType := reflect.TypeOf(item).String()
    if itemType != expectedType {
      return false
    }
  }
  return true
}

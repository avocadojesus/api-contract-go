package validators

import (
  "reflect"
  "regexp"
  "fmt"
)

func ValidateNumber(item interface{}) bool {
  itemType := reflect.TypeOf(item).String()
  return itemType == "float64"
}

func ValidateNumberCustomFormat(num interface{}, format string) bool {
  var matchFound bool
  var err interface{}
  switch(format) {
  case "int":
    matchFound, err = regexp.MatchString(`^\d{1,}$`, fmt.Sprintf("%v", num))
    break

  case "bigint":
    matchFound, err = regexp.MatchString(`^\d{1,}$`, fmt.Sprintf("%v", num))
    break

  case "float":
    matchFound, err = regexp.MatchString(`^\d{1,}\.\d{1,}$`, fmt.Sprintf("%.2f", num))
    break

  default:
    panic(fmt.Sprintf("could not validate custom string %s to format %s", num, format))
  }

  if err != nil {
    return false
  }

  return matchFound
}

func ValidateNumberArray(arr []interface{}) bool {
  for _, item := range arr {
    if !ValidateNumber(item) {
      return false
    }
  }
  return true
}

func ValidateNumberArrayCustomFormat(arr []interface{}, format string) bool {
  for _, item := range arr {
    if !ValidateNumberCustomFormat(item, format) {
      return false
    }
  }
  return true
}

func FindNumberFormat(arr []string) string {
  if SliceContains(arr, "int") {
    return "int"
  } else if SliceContains(arr, "float") {
    return "float"
  } else if SliceContains(arr, "bigint") {
    return "bigint"
  } else {
    return ""
  }
}


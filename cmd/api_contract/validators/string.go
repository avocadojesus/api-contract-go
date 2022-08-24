package validators

import (
  "regexp"
  "fmt"
)

func ValidateStringCustomFormat(str interface{}, format string) bool {
  var matchFound bool
  var err interface{}
  switch(format) {
  case "uuid":
    matchFound, err = regexp.MatchString(`^[a-zA-Z0-9]{8}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{12}$`, fmt.Sprintf("%s", str))

  case "email":
    matchFound, err = regexp.MatchString(`^.*@.*\..*$`, fmt.Sprintf("%s", str))

  case "name":
    matchFound, err = regexp.MatchString(`^[A-Za-z]*$`, fmt.Sprintf("%s", str))

  case "fullname":
    matchFound, err = regexp.MatchString(`^[A-Za-z']* [A-Za-z']*\s?[A-Za-z']{0,}$`, fmt.Sprintf("%s", str))

  default:
    panic(fmt.Sprintf("could not validate custom string %s to format %s", str, format))
  }

  if err != nil {
    return false
  }

  return matchFound
}

func ValidateStringArray(arr []interface{}) bool {
  return CheckArrayForType(arr, "string")
}

func ValidateStringArrayCustomFormat(arr []interface{}, format string) bool {
  for _, item := range arr {
    if !ValidateStringCustomFormat(item, format) {
      return false
    }
  }
  return true
}

func FindStringFormat(arr []string) string {
  if SliceContains(arr, "uuid") {
    return "uuid"
  } else if SliceContains(arr, "email") {
    return "email"
  } else if SliceContains(arr, "name") {
    return "name"
  } else if SliceContains(arr, "fullname") {
    return "fullname"
  } else {
    return ""
  }
}


package validators

import (
  "regexp"
  "fmt"
  "github.com/avocadojesus/api-contract-go/cmd/api_contract/config"
  "github.com/avocadojesus/api-contract-go/cmd/api_contract/helpers"
)

func ValidateString(
  param string,
  paramType interface{},
  results map[string]interface{},
  decorators []string,
  typeOfReturnedValue string,
  isArray bool,
  conf config.ApiContractConfig,
) bool {
  if (len(decorators) == 0) {
    if isArray {
      return validateStringArray(results[param].([]interface{}))
    } else {
      return typeOfReturnedValue == paramType
    }
  }

  format := findStringFormat(decorators)
  if format == "" {
    return false
  }

  if isArray {
    return validateStringArrayCustomFormat(results[param].([]interface{}), format)
  } else {
    return validateStringCustomFormat(results[param].(interface{}), format)
  }
}

func validateStringCustomFormat(str interface{}, format string) bool {
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
    matchFound, err = regexp.MatchString(`^[A-Za-z'.]* [A-Za-z'.]*\s?[A-Za-z'.]{0,}\s?[A-Za-z'.]{0,}$`, fmt.Sprintf("%s", str))

  default:
    panic(fmt.Sprintf("could not validate custom string %s to format %s", str, format))
  }

  if err != nil {
    return false
  }

  return matchFound
}

func validateStringArray(arr []interface{}) bool {
  return CheckArrayForType(arr, "string")
}

func validateStringArrayCustomFormat(arr []interface{}, format string) bool {
  for _, item := range arr {
    if !validateStringCustomFormat(item, format) {
      return false
    }
  }
  return true
}

func findStringFormat(arr []string) string {
  if helpers.SliceContains(arr, "uuid") {
    return "uuid"
  } else if helpers.SliceContains(arr, "email") {
    return "email"
  } else if helpers.SliceContains(arr, "name") {
    return "name"
  } else if helpers.SliceContains(arr, "fullname") {
    return "fullname"
  } else {
    return ""
  }
}


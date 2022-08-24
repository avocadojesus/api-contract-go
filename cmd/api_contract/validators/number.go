package validators

import (
  "reflect"
  "regexp"
  "fmt"
  "github.com/avocadojesus/api-contract-go/cmd/api_contract/config"
)

func ValidateNumber(
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
      return validateNumberArray(results[param].([]interface{}))
    } else {
      return validateNumberBasic(results[param].(interface{}))
    }
  }

  format := findNumberFormat(decorators)
  if format == "" {
    return false
  }

  if isArray {
    return validateNumberArrayCustomFormat(results[param].([]interface{}), format)
  } else {
    return validateNumberCustomFormat(results[param].(interface{}), format)
  }
}

func validateNumberBasic(item interface{}) bool {
  itemType := reflect.TypeOf(item).String()
  return itemType == "float64"
}

func validateNumberCustomFormat(num interface{}, format string) bool {
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

func validateNumberArray(arr []interface{}) bool {
  for _, item := range arr {
    if !validateNumberBasic(item) {
      return false
    }
  }
  return true
}

func validateNumberArrayCustomFormat(arr []interface{}, format string) bool {
  for _, item := range arr {
    if !validateNumberCustomFormat(item, format) {
      return false
    }
  }
  return true
}

func findNumberFormat(arr []string) string {
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


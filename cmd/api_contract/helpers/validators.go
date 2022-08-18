package api_contract_helpers

import (
  "reflect"
  "time"
  "fmt"
)

const DATE_FORMAT = "2006-01-02"
const DATETIME_FORMAT = time.RFC3339

func ValidatePayloadShapeForParam(param string, paramType interface{}, results map[string]interface{}) bool {
  if results[param] == nil {
    return false
  }

  // when using nested data structures, just check that they both are type map,
  // since each of their inner children will be tested separately
  if IsMap(results[param]) && IsMap(paramType) {
    for key, _ := range results[param].(map[string]interface{}) {
      expectedType := paramType.(map[string]interface{})[key]
      if !ValidatePayloadShapeForParam(key, expectedType, results[param].(map[string]interface{})) {
        return false
      }
    }
    return true

  } else {
    typeOfReturnedValue := reflect.TypeOf(results[param]).String()

    switch paramType {
    case "bool[]":
      return ValidateBoolArray(results[param].([]interface{}))

    case "date":
      return ValidateDate(results[param].(interface{}))

    case "date[]":
      return ValidateDateArray(results[param].([]interface{}))

    case "datetime":
      return ValidateDatetime(results[param].(interface{}))

    case "datetime[]":
      return ValidateDatetimeArray(results[param].([]interface{}))

    case "number":
      return ValidateNumber(results[param].(interface{}))

    case "number[]":
      return ValidateNumberArray(results[param].([]interface{}))

    case "string[]":
      return ValidateStringArray(results[param].([]interface{}))

    default:
      return typeOfReturnedValue == paramType
    }
  }
}

func CheckPayloadForUnexpectedKeys(results map[string]interface{}, endpoints map[string]interface{}) bool {
  for key, val := range results {
    if IsMap(val) && IsMap(endpoints[key]) {
      return CheckPayloadForUnexpectedKeys(results[key].(map[string]interface{}), endpoints[key].(map[string]interface{}))
    } else {
      if results[key] != nil && endpoints[key] == nil {
        return false
      }
    }
  }
  return true
}

func ValidateDate(date interface{}) bool {
  _, err := time.Parse(DATE_FORMAT, fmt.Sprintf("%s", date))
  return err == nil
}

func ValidateDatetime(date interface{}) bool {
  _, err := time.Parse(DATETIME_FORMAT, fmt.Sprintf("%s", date))
  fmt.Println(DATETIME_FORMAT, fmt.Sprintf("%s", date), err)
  return err == nil
}

func ValidateNumber(item interface{}) bool {
  itemType := reflect.TypeOf(item).String()
  return itemType == "float64"
}

func ValidateBoolArray(arr []interface{}) bool {
  return checkArrayForType(arr, "bool")
}

func ValidateStringArray(arr []interface{}) bool {
  return checkArrayForType(arr, "string")
}

func ValidateNumberArray(arr []interface{}) bool {
  for _, item := range arr {
    if !ValidateNumber(item) {
      return false
    }
  }
  return true
}

func ValidateDateArray(arr []interface{}) bool {
  for _, item := range arr {
    if !ValidateDate(item) {
      return false
    }
  }
  return true
}

func ValidateDatetimeArray(arr []interface{}) bool {
  for _, item := range arr {
    if !ValidateDatetime(item) {
      return false
    }
  }
  return true
}

func checkArrayForType(arr []interface{}, expectedType string) bool {
  for _, item := range arr {
    itemType := reflect.TypeOf(item).String()
    if itemType != expectedType {
      return false
    }
  }
  return true
}


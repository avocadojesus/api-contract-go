package helpers

import (
  "reflect"
  "time"
  "fmt"
  "regexp"
)

const DATE_FORMAT = "2006-01-02"
const DATETIME_FORMAT = time.RFC3339

func ValidateParam(param string, paramType interface{}, results map[string]interface{}) bool {
  if results[param] == nil {
    return false
  }

  // when using nested data structures, just check that they both are type map,
  // since each of their inner children will be tested separately
  if IsMap(results[param]) && IsMap(paramType) {
    for key, _ := range results[param].(map[string]interface{}) {
      expectedType := paramType.(map[string]interface{})[key]
      if !ValidateParam(key, expectedType, results[param].(map[string]interface{})) {
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
      datatype, decorators, isArray := ParseDatatype(fmt.Sprintf("%s", paramType))

      switch datatype {
      case "date":
        format := findDateFormat(decorators)
        if format == "" {
          return false
        }

        if isArray {
          return ValidateDateArrayCustomFormat(results[param].([]interface{}), format)
        } else {
          return ValidateDateCustomFormat(results[param].(interface{}), format)
        }

      case "datetime":
        format := findDatetimeFormat(decorators)
        if format == "" {
          return false
        }

        if isArray {
          return ValidateDatetimeArrayCustomFormat(results[param].([]interface{}), format)
        } else {
          return ValidateDatetimeCustomFormat(results[param].(interface{}), format)
        }

      default:
        return typeOfReturnedValue == paramType
      }
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
  return err == nil
}

func ValidateDatetimeCustomFormat(date interface{}, format string) bool {
  _, err := time.Parse(format, fmt.Sprintf("%s", date))
  return err == nil
}

func ValidateDateCustomFormat(date interface{}, format string) bool {
  matchFound, err := regexp.MatchString(format, fmt.Sprintf("%s", date))
  if err != nil {
    return false
  }
  return matchFound
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

func ValidateDateArrayCustomFormat(arr []interface{}, format string) bool {
  for _, item := range arr {
    if !ValidateDateCustomFormat(item, format) {
      return false
    }
  }
  return true
}

func ValidateDatetimeArrayCustomFormat(arr []interface{}, format string) bool {
  for _, item := range arr {
    if !ValidateDatetimeCustomFormat(item, format) {
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

func findDateFormat(arr []string) string {
  if SliceContains(arr, "mmddyyyy") || SliceContains(arr, "MMDDYYYY") {
    return Date.MMDDYYYY
  } else if SliceContains(arr, "mmddyy") || SliceContains(arr, "MMDDYY") {
    return Date.MMDDYY
  } else if SliceContains(arr, "yyyymmdd") || SliceContains(arr, "YYYYMMDD") {
    return Date.YYYYMMDD
  } else if SliceContains(arr, "yymmdd") || SliceContains(arr, "YYMMDD") {
    return Date.YYMMDD
  } else {
    return ""
  }
}

func findDatetimeFormat(arr []string) string {
  if SliceContains(arr, "ansic") {
    return time.ANSIC
  } else if (SliceContains(arr, "unix_date") || SliceContains(arr, "unix")) {
    return time.UnixDate
  } else if (SliceContains(arr, "ruby_date") || SliceContains(arr, "ruby")) {
    return time.RubyDate
  } else if (SliceContains(arr, "rfc822") || SliceContains(arr, "RFC822")) {
    return time.RFC822
  } else if (SliceContains(arr, "rfc822z") || SliceContains(arr, "RFC822Z")) {
    return time.RFC822Z
  } else if (SliceContains(arr, "rfc850") || SliceContains(arr, "RFC850")) {
    return time.RFC850
  } else if (SliceContains(arr, "rfc1123") || SliceContains(arr, "RFC1123")) {
    return time.RFC1123
  } else if (SliceContains(arr, "rfc1123z") || SliceContains(arr, "RFC1123Z")) {
    return time.RFC1123Z
  } else if (SliceContains(arr, "rfc3339") || SliceContains(arr, "RFC3339")) {
    return time.RFC3339
  } else if (SliceContains(arr, "rfc3339_nano") || SliceContains(arr, "RFC3339Nano")) {
    return time.RFC3339Nano
  } else if (SliceContains(arr, "kitchen") || SliceContains(arr, "Kitchen")) {
    return time.Kitchen
  } else if
    SliceContains(arr, "stamp") ||
      SliceContains(arr, "Stamp") ||
      SliceContains(arr, "timestamp") ||
      SliceContains(arr, "Timestamp") {
    return time.Stamp
  } else if
    SliceContains(arr, "stamp_milli") ||
      SliceContains(arr, "StampMilli") ||
      SliceContains(arr, "timestamp_milli") ||
      SliceContains(arr, "TimestampMilli") {
    return time.StampMilli
  } else if
    SliceContains(arr, "stamp_micro") ||
      SliceContains(arr, "StampMicro") ||
      SliceContains(arr, "timestamp_micro") ||
      SliceContains(arr, "TimestampMicro") {
    return time.StampMicro
  } else if
    SliceContains(arr, "stamp_nano") ||
      SliceContains(arr, "StampNano") ||
      SliceContains(arr, "timestamp_nano") ||
      SliceContains(arr, "TimestampNano") {
    return time.StampNano
  } else {
    return ""
  }
}


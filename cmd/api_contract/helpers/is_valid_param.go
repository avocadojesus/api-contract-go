package helpers

import (
  "reflect"
  "time"
  "fmt"
  "regexp"
)

const DATE_FORMAT = "2006-01-02"
const DATETIME_FORMAT = time.RFC3339

func IsValidParam(param string, paramType interface{}, results map[string]interface{}) bool {
  if results[param] == nil {
    return false
  }

  // when using nested data structures, just check that they both are type map,
  // since each of their inner children will be tested separately
  if IsMap(results[param]) && IsMap(paramType) {
    for key, _ := range results[param].(map[string]interface{}) {
      expectedType := paramType.(map[string]interface{})[key]
      if !IsValidParam(key, expectedType, results[param].(map[string]interface{})) {
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
      return validateDate(results[param].(interface{}))

    case "date[]":
      return validateDateArray(results[param].([]interface{}))

    case "datetime":
      return validateDatetime(results[param].(interface{}))

    case "datetime[]":
      return validateDatetimeArray(results[param].([]interface{}))

    case "number":
      return validateNumber(results[param].(interface{}))

    case "number[]":
      return validateNumberArray(results[param].([]interface{}))

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
          return validateDateArrayCustomFormat(results[param].([]interface{}), format)
        } else {
          return validateDateCustomFormat(results[param].(interface{}), format)
        }

      case "datetime":
        format := findDatetimeFormat(decorators)
        if format == "" {
          return false
        }

        if isArray {
          return validateDatetimeArrayCustomFormat(results[param].([]interface{}), format)
        } else {
          return validateDatetimeCustomFormat(results[param].(interface{}), format)
        }

      default:
        return typeOfReturnedValue == paramType
      }
    }
  }
}

func validateDate(date interface{}) bool {
  _, err := time.Parse(DATE_FORMAT, fmt.Sprintf("%s", date))
  return err == nil
}

func validateDatetime(date interface{}) bool {
  _, err := time.Parse(DATETIME_FORMAT, fmt.Sprintf("%s", date))
  return err == nil
}

func validateDatetimeCustomFormat(date interface{}, format string) bool {
  _, err := time.Parse(format, fmt.Sprintf("%s", date))
  return err == nil
}

func validateDateCustomFormat(date interface{}, format string) bool {
  matchFound, err := regexp.MatchString(format, fmt.Sprintf("%s", date))
  if err != nil {
    return false
  }
  return matchFound
}

func validateNumber(item interface{}) bool {
  itemType := reflect.TypeOf(item).String()
  return itemType == "float64"
}

func ValidateBoolArray(arr []interface{}) bool {
  return checkArrayForType(arr, "bool")
}

func ValidateStringArray(arr []interface{}) bool {
  return checkArrayForType(arr, "string")
}

func validateNumberArray(arr []interface{}) bool {
  for _, item := range arr {
    if !validateNumber(item) {
      return false
    }
  }
  return true
}

func validateDateArray(arr []interface{}) bool {
  for _, item := range arr {
    if !validateDate(item) {
      return false
    }
  }
  return true
}

func validateDatetimeArray(arr []interface{}) bool {
  for _, item := range arr {
    if !validateDatetime(item) {
      return false
    }
  }
  return true
}

func validateDateArrayCustomFormat(arr []interface{}, format string) bool {
  for _, item := range arr {
    if !validateDateCustomFormat(item, format) {
      return false
    }
  }
  return true
}

func validateDatetimeArrayCustomFormat(arr []interface{}, format string) bool {
  for _, item := range arr {
    if !validateDatetimeCustomFormat(item, format) {
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

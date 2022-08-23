package helpers

import (
  "reflect"
  "time"
  "fmt"
  "regexp"
  "github.com/avocadojesus/api-contract-go/cmd/api_contract/validators"
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
      return validators.ValidateBoolArray(results[param].([]interface{}))

    case "date":
      return validateDate(results[param].(interface{}))

    case "date[]":
      return validateDateArray(results[param].([]interface{}))

    case "datetime":
      return validateDatetime(results[param].(interface{}))

    case "datetime[]":
      return validateDatetimeArray(results[param].([]interface{}))

    case "number":
      return validators.ValidateNumber(results[param].(interface{}))

    case "number[]":
      return validators.ValidateNumberArray(results[param].([]interface{}))

    case "string[]":
      return validators.ValidateStringArray(results[param].([]interface{}))

    default:
      datatype, decorators, isArray := ParseDatatype(fmt.Sprintf("%s", paramType))

      switch datatype {
      case "string":
        if (len(decorators) == 0) {
          return typeOfReturnedValue == paramType
        }

        format := validators.FindStringFormat(decorators)
        if format == "" {
          return false
        }

        if isArray {
          return validators.ValidateStringArrayCustomFormat(results[param].([]interface{}), format)
        } else {
          return validators.ValidateStringCustomFormat(results[param].(interface{}), format)
        }

      case "number":
        if (len(decorators) == 0) {
          return validators.ValidateNumber(results[param].(interface{}))
        }

        format := validators.FindNumberFormat(decorators)
        if format == "" {
          return false
        }

        if isArray {
          return validators.ValidateNumberArrayCustomFormat(results[param].([]interface{}), format)
        } else {
          return validators.ValidateNumberCustomFormat(results[param].(interface{}), format)
        }

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

func findDateFormat(arr []string) string {
  if validators.SliceContains(arr, "mmddyyyy") || validators.SliceContains(arr, "MMDDYYYY") {
    return Date.MMDDYYYY
  } else if validators.SliceContains(arr, "mmddyy") || validators.SliceContains(arr, "MMDDYY") {
    return Date.MMDDYY
  } else if validators.SliceContains(arr, "yyyymmdd") || validators.SliceContains(arr, "YYYYMMDD") {
    return Date.YYYYMMDD
  } else if validators.SliceContains(arr, "yymmdd") || validators.SliceContains(arr, "YYMMDD") {
    return Date.YYMMDD
  } else {
    return ""
  }
}

func findDatetimeFormat(arr []string) string {
  if validators.SliceContains(arr, "ansic") {
    return time.ANSIC
  } else if (validators.SliceContains(arr, "iso861") || validators.SliceContains(arr, "ISO861")) {
    return time.RFC3339
  } else if (validators.SliceContains(arr, "unix_date") || validators.SliceContains(arr, "unix")) {
    return time.UnixDate
  } else if (validators.SliceContains(arr, "ruby_date") || validators.SliceContains(arr, "ruby")) {
    return time.RubyDate
  } else if (validators.SliceContains(arr, "rfc822") || validators.SliceContains(arr, "RFC822")) {
    return time.RFC822
  } else if (validators.SliceContains(arr, "rfc822z") || validators.SliceContains(arr, "RFC822Z")) {
    return time.RFC822Z
  } else if (validators.SliceContains(arr, "rfc850") || validators.SliceContains(arr, "RFC850")) {
    return time.RFC850
  } else if (validators.SliceContains(arr, "rfc1123") || validators.SliceContains(arr, "RFC1123")) {
    return time.RFC1123
  } else if (validators.SliceContains(arr, "rfc1123z") || validators.SliceContains(arr, "RFC1123Z")) {
    return time.RFC1123Z
  } else if (validators.SliceContains(arr, "rfc3339") || validators.SliceContains(arr, "RFC3339")) {
    return time.RFC3339
  } else if (validators.SliceContains(arr, "rfc3339_nano") || validators.SliceContains(arr, "RFC3339Nano")) {
    return time.RFC3339Nano
  } else if (validators.SliceContains(arr, "kitchen") || validators.SliceContains(arr, "Kitchen")) {
    return time.Kitchen
  } else if
    validators.SliceContains(arr, "stamp") ||
      validators.SliceContains(arr, "Stamp") ||
      validators.SliceContains(arr, "timestamp") ||
      validators.SliceContains(arr, "Timestamp") {
    return time.Stamp
  } else if
    validators.SliceContains(arr, "stamp_milli") ||
      validators.SliceContains(arr, "StampMilli") ||
      validators.SliceContains(arr, "timestamp_milli") ||
      validators.SliceContains(arr, "TimestampMilli") {
    return time.StampMilli
  } else if
    validators.SliceContains(arr, "stamp_micro") ||
      validators.SliceContains(arr, "StampMicro") ||
      validators.SliceContains(arr, "timestamp_micro") ||
      validators.SliceContains(arr, "TimestampMicro") {
    return time.StampMicro
  } else if
    validators.SliceContains(arr, "stamp_nano") ||
      validators.SliceContains(arr, "StampNano") ||
      validators.SliceContains(arr, "timestamp_nano") ||
      validators.SliceContains(arr, "TimestampNano") {
    return time.StampNano
  } else {
    return ""
  }
}


package validators

import (
  "reflect"
  "time"
  "fmt"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract/helpers"
)

const DATE_FORMAT = "2006-01-02"
const DATETIME_FORMAT = time.RFC3339

func IsValidParam(param string, paramType interface{}, results map[string]interface{}) bool {
  if results[param] == nil {
    return false
  }

  // when using nested data structures, just check that they both are type map,
  // since each of their inner children will be tested separately
  if helpers.IsMap(results[param]) && helpers.IsMap(paramType) {
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
      return ValidateDate(results[param].(interface{}))

    case "date[]":
      return ValidateDateArray(results[param].([]interface{}))

    case "datetime":
      return validateDatetime(results[param].(interface{}))

    case "datetime[]":
      return validateDatetimeArray(results[param].([]interface{}))

    case "number":
      return ValidateNumber(results[param].(interface{}))

    case "number[]":
      return ValidateNumberArray(results[param].([]interface{}))

    case "string[]":
      return ValidateStringArray(results[param].([]interface{}))

    default:
      datatype, decorators, isArray := helpers.ParseDatatype(fmt.Sprintf("%s", paramType))

      switch datatype {
      case "string":
        if (len(decorators) == 0) {
          return typeOfReturnedValue == paramType
        }

        format := FindStringFormat(decorators)
        if format == "" {
          return false
        }

        if isArray {
          return ValidateStringArrayCustomFormat(results[param].([]interface{}), format)
        } else {
          return ValidateStringCustomFormat(results[param].(interface{}), format)
        }

      case "number":
        if (len(decorators) == 0) {
          return ValidateNumber(results[param].(interface{}))
        }

        format := FindNumberFormat(decorators)
        if format == "" {
          return false
        }

        if isArray {
          return ValidateNumberArrayCustomFormat(results[param].([]interface{}), format)
        } else {
          return ValidateNumberCustomFormat(results[param].(interface{}), format)
        }

      case "date":
        format := FindDateFormat(decorators)
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

func validateDatetime(date interface{}) bool {
  _, err := time.Parse(DATETIME_FORMAT, fmt.Sprintf("%s", date))
  return err == nil
}

func validateDatetimeCustomFormat(date interface{}, format string) bool {
  _, err := time.Parse(format, fmt.Sprintf("%s", date))
  return err == nil
}

func validateDatetimeArray(arr []interface{}) bool {
  for _, item := range arr {
    if !validateDatetime(item) {
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

func findDatetimeFormat(arr []string) string {
  if SliceContains(arr, "ansic") {
    return time.ANSIC
  } else if (SliceContains(arr, "iso861") || SliceContains(arr, "ISO861")) {
    return time.RFC3339
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


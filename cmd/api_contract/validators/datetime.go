package validators

import (
  "time"
  "fmt"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract/config"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract/helpers"
)

func ValidateDatetime(
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
      return validateDatetimeArray(results[param].([]interface{}))
    } else {
      return validateDatetimeBasic(results[param].(interface{}))
    }
  }

  format := findDatetimeFormat(decorators)
  if format == "" {
    return false
  }

  if isArray {
    return validateDatetimeArrayCustomFormat(results[param].([]interface{}), format)
  } else {
    return validateDatetimeCustomFormat(results[param].(interface{}), format)
  }
}

func validateDatetimeBasic(date interface{}) bool {
  _, err := time.Parse(DATETIME_FORMAT, fmt.Sprintf("%s", date))
  return err == nil
}

func validateDatetimeCustomFormat(date interface{}, format string) bool {
  _, err := time.Parse(format, fmt.Sprintf("%s", date))
  return err == nil
}

func validateDatetimeArray(arr []interface{}) bool {
  for _, item := range arr {
    if !validateDatetimeBasic(item) {
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
  if helpers.SliceContains(arr, "ansic") {
    return time.ANSIC
  } else if (helpers.SliceContains(arr, "iso861") || helpers.SliceContains(arr, "ISO861")) {
    return time.RFC3339
  } else if (helpers.SliceContains(arr, "unix_date") || helpers.SliceContains(arr, "unix")) {
    return time.UnixDate
  } else if (helpers.SliceContains(arr, "ruby_date") || helpers.SliceContains(arr, "ruby")) {
    return time.RubyDate
  } else if (helpers.SliceContains(arr, "rfc822") || helpers.SliceContains(arr, "RFC822")) {
    return time.RFC822
  } else if (helpers.SliceContains(arr, "rfc822z") || helpers.SliceContains(arr, "RFC822Z")) {
    return time.RFC822Z
  } else if (helpers.SliceContains(arr, "rfc850") || helpers.SliceContains(arr, "RFC850")) {
    return time.RFC850
  } else if (helpers.SliceContains(arr, "rfc1123") || helpers.SliceContains(arr, "RFC1123")) {
    return time.RFC1123
  } else if (helpers.SliceContains(arr, "rfc1123z") || helpers.SliceContains(arr, "RFC1123Z")) {
    return time.RFC1123Z
  } else if (helpers.SliceContains(arr, "rfc3339") || helpers.SliceContains(arr, "RFC3339")) {
    return time.RFC3339
  } else if (helpers.SliceContains(arr, "rfc3339_nano") || helpers.SliceContains(arr, "RFC3339Nano")) {
    return time.RFC3339Nano
  } else if (helpers.SliceContains(arr, "kitchen") || helpers.SliceContains(arr, "Kitchen")) {
    return time.Kitchen
  } else if
    helpers.SliceContains(arr, "stamp") ||
      helpers.SliceContains(arr, "Stamp") ||
      helpers.SliceContains(arr, "timestamp") ||
      helpers.SliceContains(arr, "Timestamp") {
    return time.Stamp
  } else if
    helpers.SliceContains(arr, "stamp_milli") ||
      helpers.SliceContains(arr, "StampMilli") ||
      helpers.SliceContains(arr, "timestamp_milli") ||
      helpers.SliceContains(arr, "TimestampMilli") {
    return time.StampMilli
  } else if
    helpers.SliceContains(arr, "stamp_micro") ||
      helpers.SliceContains(arr, "StampMicro") ||
      helpers.SliceContains(arr, "timestamp_micro") ||
      helpers.SliceContains(arr, "TimestampMicro") {
    return time.StampMicro
  } else if
    helpers.SliceContains(arr, "stamp_nano") ||
      helpers.SliceContains(arr, "StampNano") ||
      helpers.SliceContains(arr, "timestamp_nano") ||
      helpers.SliceContains(arr, "TimestampNano") {
    return time.StampNano
  } else {
    return ""
  }
}


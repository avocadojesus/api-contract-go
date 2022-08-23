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
        format := FindDatetimeFormat(decorators)
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

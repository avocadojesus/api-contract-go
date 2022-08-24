package validators

import (
  "fmt"
  "reflect"
  "github.com/avocadojesus/api-contract-go/cmd/api_contract/config"
  "github.com/avocadojesus/api-contract-go/cmd/api_contract/helpers"
)

func ValidateParam(
  param string,
  paramType interface{},
  results map[string]interface{},
  conf config.ApiContractConfig,
) bool {
  if results[param] == nil {
    return false
  }

  // when using nested data structures, just check that they both are type map,
  // since each of their inner children will be tested separately
  if helpers.IsMap(results[param]) && helpers.IsMap(paramType) {
    for key, _ := range results[param].(map[string]interface{}) {
      expectedType := paramType.(map[string]interface{})[key]
      if !ValidateParam(key, expectedType, results[param].(map[string]interface{}), conf) {
        return false
      }
    }
    return true

  } else {
    typeOfReturnedValue := reflect.TypeOf(results[param]).String()
    datatype, decorators, isArray := helpers.ParseDatatype(fmt.Sprintf("%s", paramType))

    switch datatype {
    case "string":
      return ValidateString(param, paramType, results, decorators, typeOfReturnedValue, isArray, conf)

    case "bool":
      return ValidateBool(param, paramType, results, typeOfReturnedValue, isArray)

    case "number":
      return ValidateNumber(param, paramType, results, decorators, typeOfReturnedValue, isArray, conf)

    case "date":
      return ValidateDate(param, paramType, results, decorators, typeOfReturnedValue, isArray, conf)

    case "datetime":
      return ValidateDatetime(param, paramType, results, decorators, typeOfReturnedValue, isArray, conf)

    default:
      if conf.Serializers[datatype] != nil {
        return ValidateFromSerializer(conf, datatype, results, param, isArray)
      } else {
        return typeOfReturnedValue == paramType
      }
    }
  }
}

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
      return validateString(param, paramType, results, decorators, typeOfReturnedValue, isArray, conf)

    case "bool":
      return validateBool(param, paramType, results, typeOfReturnedValue, isArray)

    case "number":
      return validateNumber(param, paramType, results, decorators, typeOfReturnedValue, isArray, conf)

    case "date":
      return validateDate(param, paramType, results, decorators, typeOfReturnedValue, isArray, conf)

    case "datetime":
      return validateDatetime(param, paramType, results, decorators, typeOfReturnedValue, isArray, conf)

    default:
      if conf.Serializers[datatype] != nil {
        return ValidateFromSerializer(conf, datatype, results, param, isArray)
      } else {
        return typeOfReturnedValue == paramType
      }
    }
  }
}

func validateDatetime(
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
      return ValidateDatetimeArray(results[param].([]interface{}))
    } else {
      return ValidateDatetime(results[param].(interface{}))
    }
  }

  format := FindDatetimeFormat(decorators)
  if format == "" {
    return false
  }

  if isArray {
    return ValidateDatetimeArrayCustomFormat(results[param].([]interface{}), format)
  } else {
    return ValidateDatetimeCustomFormat(results[param].(interface{}), format)
  }
}

func validateDate(
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
      return ValidateDateArray(results[param].([]interface{}))
    } else {
      return ValidateDate(results[param].(interface{}))
    }
  }

  format := FindDateFormat(decorators)
  if format == "" {
    return false
  }

  if isArray {
    return ValidateDateArrayCustomFormat(results[param].([]interface{}), format)
  } else {
    return ValidateDateCustomFormat(results[param].(interface{}), format)
  }
}

func validateNumber(
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
      return ValidateNumberArray(results[param].([]interface{}))
    } else {
      return ValidateNumber(results[param].(interface{}))
    }
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
}

func validateBool(
  param string,
  paramType interface{},
  results map[string]interface{},
  typeOfReturnedValue string,
  isArray bool,
) bool {
  if isArray {
    return ValidateBoolArray(results[param].([]interface{}))
  } else {
    return typeOfReturnedValue == paramType
  }
}

func validateString(
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
      return ValidateStringArray(results[param].([]interface{}))
    } else {
      return typeOfReturnedValue == paramType
    }
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
}

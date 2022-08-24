package api_contract

import (
  "fmt"
  "encoding/json"
  "reflect"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract/helpers"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract/validators"
)

type ApiContractConfig struct {
  Serializers map[string]map[string]interface{} `json:"serializers"`
}

type EndpointData struct {
  PayloadShape map[string]interface{} `json:"payload_shape"`
}

type ApiContractJSON struct {
  Config ApiContractConfig `json:"config"`
}

func Validate(bytes []byte, httpMethod string, endpoint string) (bool, string) {
  var results map[string]interface{}
  json.Unmarshal(bytes, &results)

  endpoints, config := unmarshalJSONToEndpointData("./api-contract.json")
  endpointKey := httpMethod + ":" + endpoint
  foundEndpoint := endpoints[endpointKey]

  if foundEndpoint != nil {
    for param, paramType := range foundEndpoint.PayloadShape {
      if !ValidateParam(param, paramType, results, config) {
        return false, fmt.Sprintf("The param `%s` does not match expected type `%s`", param, paramType)
      }
    }
  }

  if !helpers.HasUnrecognizedParams(results, foundEndpoint.PayloadShape) {
    return false, "Unexpected keys found"
  }

  return true, ""
}

func unmarshalJSONToEndpointData(path string) (map[string]*EndpointData, ApiContractConfig) {
  data := ReadJSON(path)
	endpoints := map[string]*EndpointData{}
  if err := json.Unmarshal(data, &endpoints); err != nil {
    panic(err)
  }


  var jsonData ApiContractJSON
  var config ApiContractConfig
  if err2 := json.Unmarshal(data, &jsonData); err2 != nil {
    config = ApiContractConfig{}
  } else {
    config = jsonData.Config
  }

  return endpoints, config
}

func ValidateParam(param string, paramType interface{}, results map[string]interface{}, config ApiContractConfig) bool {
  if results[param] == nil {
    return false
  }

  // when using nested data structures, just check that they both are type map,
  // since each of their inner children will be tested separately
  if helpers.IsMap(results[param]) && helpers.IsMap(paramType) {
    for key, _ := range results[param].(map[string]interface{}) {
      expectedType := paramType.(map[string]interface{})[key]
      if !ValidateParam(key, expectedType, results[param].(map[string]interface{}), config) {
        return false
      }
    }
    return true

  } else {
    typeOfReturnedValue := reflect.TypeOf(results[param]).String()
    datatype, decorators, isArray := helpers.ParseDatatype(fmt.Sprintf("%s", paramType))

    switch datatype {
    case "string":
      if (len(decorators) == 0) {
        if isArray {
          return validators.ValidateStringArray(results[param].([]interface{}))
        } else {
          return typeOfReturnedValue == paramType
        }
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

    case "bool":
      if isArray {
        return validators.ValidateBoolArray(results[param].([]interface{}))
      } else {
        return typeOfReturnedValue == paramType
      }

    case "number":
      if (len(decorators) == 0) {
        if isArray {
          return validators.ValidateNumberArray(results[param].([]interface{}))
        } else {
          return validators.ValidateNumber(results[param].(interface{}))
        }
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
      if (len(decorators) == 0) {
        if isArray {
          return validators.ValidateDateArray(results[param].([]interface{}))
        } else {
          return validators.ValidateDate(results[param].(interface{}))
        }
      }

      format := validators.FindDateFormat(decorators)
      if format == "" {
        return false
      }

      if isArray {
        return validators.ValidateDateArrayCustomFormat(results[param].([]interface{}), format)
      } else {
        return validators.ValidateDateCustomFormat(results[param].(interface{}), format)
      }

    case "datetime":
      if (len(decorators) == 0) {
        if isArray {
          return validators.ValidateDatetimeArray(results[param].([]interface{}))
        } else {
          return validators.ValidateDatetime(results[param].(interface{}))
        }
      }

      format := validators.FindDatetimeFormat(decorators)
      if format == "" {
        return false
      }

      if isArray {
        return validators.ValidateDatetimeArrayCustomFormat(results[param].([]interface{}), format)
      } else {
        return validators.ValidateDatetimeCustomFormat(results[param].(interface{}), format)
      }

    default:
      if config.Serializers[datatype] != nil {
        return validateFromSerializer(config, datatype, results, param, isArray)
      } else {
        return typeOfReturnedValue == paramType
      }
    }
  }
}

func validateFromSerializer(
  config ApiContractConfig,
  serializerName string,
  results map[string]interface{},
  param string,
  isArray bool,
) bool {
  serializer := config.Serializers[serializerName]
  for serializerParam, serializerParamType := range serializer {
    if isArray {
      obj := results[param].([]interface{})
      for _, resultObj := range obj {
        if !ValidateParam(serializerParam, serializerParamType, resultObj.(map[string]interface{}), config) {
          return false
        }
      }
    } else {
      obj := results[param].(map[string]interface{})
      if !ValidateParam(serializerParam, serializerParamType, obj, config) {
        return false
      }
    }
  }
  return true
}

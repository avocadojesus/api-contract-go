package validators

import (
  "github.com/avocadojesus/api-contract-go/cmd/api_contract/config"
)

func ValidateFromSerializer(
  config config.ApiContractConfig,
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

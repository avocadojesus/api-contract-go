package api_contract

import (
  "fmt"
  "encoding/json"
  "github.com/avocadojesus/api-contract-go/cmd/api_contract/config"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract/helpers"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract/validators"
)

func Validate(bytes []byte, httpMethod string, endpoint string) (bool, string) {
  var results map[string]interface{}
  json.Unmarshal(bytes, &results)

  endpoints, config := parseApiContractJSON("./api-contract.json")
  endpointKey := httpMethod + ":" + endpoint
  foundEndpoint := endpoints[endpointKey]

  if foundEndpoint != nil {
    for param, paramType := range foundEndpoint.PayloadShape {
      if !validators.ValidateParam(param, paramType, results, config) {
        return false, fmt.Sprintf("The param `%s` does not match expected type `%s`", param, paramType)
      }
    }
  }

  if !helpers.HasUnrecognizedParams(results, foundEndpoint.PayloadShape) {
    return false, "Unexpected keys found"
  }

  return true, ""
}

func parseApiContractJSON(path string) (map[string]*config.EndpointData, config.ApiContractConfig) {
  data := ReadJSON(path)
	endpoints := map[string]*config.EndpointData{}
  if err := json.Unmarshal(data, &endpoints); err != nil {
    panic(err)
  }

  var jsonData config.ApiContractJSON
  var conf config.ApiContractConfig
  if err2 := json.Unmarshal(data, &jsonData); err2 != nil {
    conf = config.ApiContractConfig{}
  } else {
    conf = jsonData.Config
  }

  return endpoints, conf
}


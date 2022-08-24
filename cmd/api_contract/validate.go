package api_contract

import (
  "fmt"
  "encoding/json"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract/helpers"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract/validators"
)

func Validate(bytes []byte, httpMethod string, endpoint string) (bool, string) {
  var results map[string]interface{}
  json.Unmarshal(bytes, &results)

  endpoints, config := unmarshalJSONToEndpointData("./api-contract.json")
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

type EndpointData struct {
  PayloadShape map[string]interface{} `json:"payload_shape"`
}

type ApiContractJSON struct {
  Config validators.ApiContractConfig `json:"config"`
}

func unmarshalJSONToEndpointData(path string) (map[string]*EndpointData, validators.ApiContractConfig) {
  data := ReadJSON(path)
	endpoints := map[string]*EndpointData{}
  if err := json.Unmarshal(data, &endpoints); err != nil {
    panic(err)
  }


  var jsonData ApiContractJSON
  var config validators.ApiContractConfig
  if err2 := json.Unmarshal(data, &jsonData); err2 != nil {
    config = validators.ApiContractConfig{}
  } else {
    config = jsonData.Config
  }

  return endpoints, config
}

package api_contract

import (
  "fmt"
  "encoding/json"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract/helpers"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract/validators"
)

type EndpointData struct {
  PayloadShape map[string]interface{} `json:"payload_shape"`
}

func Validate(bytes []byte, httpMethod string, endpoint string) (bool, string) {
  var results map[string]interface{}
  json.Unmarshal(bytes, &results)

  endpoints := unmarshalJSONToEndpointData("./api-contract.json")
  endpointKey := httpMethod + ":" + endpoint
  foundEndpoint := endpoints[endpointKey]

  if foundEndpoint != nil {
    for param, paramType := range foundEndpoint.PayloadShape {
      if !validators.IsValidParam(param, paramType, results) {
        return false, fmt.Sprintf("The param `%s` does not match expected type `%s`", param, paramType)
      }
    }
  }

  if !helpers.HasUnrecognizedParams(results, foundEndpoint.PayloadShape) {
    return false, "Unexpected keys found"
  }

  return true, ""
}

func unmarshalJSONToEndpointData(path string) map[string]*EndpointData {
  data := ReadJSON(path)
	endpoints := map[string]*EndpointData{}
  if err := json.Unmarshal(data, &endpoints); err != nil {
    panic(err)
  }
  return endpoints
}

package api_contract

import (
  "fmt"
  "encoding/json"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract/helpers"
)

type EndpointData struct {
  PayloadShape map[string]interface{} `json:"payload_shape"`
}

func Validate(bytes []byte, httpMethod string, endpoint string) (bool, string) {
  var results map[string]interface{}
  json.Unmarshal(bytes, &results)

  endpoints := unmarshalJSONToEndpointData("./endpoints.json")
  endpointKey := httpMethod + ":" + endpoint
  foundEndpoint := endpoints[endpointKey]

  if foundEndpoint != nil {
    for param, paramType := range foundEndpoint.PayloadShape {
      if !api_contract_helpers.ValidateParam(param, paramType, results) {
        return false, fmt.Sprintf("The param `%s` does not match expected type `%s`", param, paramType)
      }
    }
  }

  if !api_contract_helpers.CheckPayloadForUnexpectedKeys(results, foundEndpoint.PayloadShape) {
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

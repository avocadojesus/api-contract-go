package config

type ApiContractConfig struct {
  Serializers map[string]map[string]interface{} `json:"serializers"`
}

type EndpointData struct {
  PayloadShape map[string]interface{} `json:"payload_shape"`
}

type ApiContractJSON struct {
  Config ApiContractConfig `json:"config"`
}

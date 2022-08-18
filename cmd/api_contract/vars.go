package api_contract

import (
	"api-contract-go/cmd/api_contract/helpers"
)

// this is used for stubbing in tests
var ReadJSON = api_contract_helpers.ReadJSON
var InternalValidatePayload = ValidatePayload

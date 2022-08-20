package api_contract

import (
	"github.com/avocadojesus/api-contract-go/cmd/api_contract/helpers"
)

// this is used for stubbing in tests
var ReadJSON = helpers.ReadJSON
var InternalValidate = Validate

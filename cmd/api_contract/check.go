package api_contract

import (
  "fmt"
)

func Check(bytes []byte, httpMethod string, endpoint string, err func(args ...any)) {
  passedValidation, reason := InternalValidate(bytes, httpMethod, endpoint)

  if !passedValidation {
    err(fmt.Sprintf("JSON shape did not match expected. Reason given: %s", reason))
  }
}

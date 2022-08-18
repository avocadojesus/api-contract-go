package api_contract_test

import (
  "testing"
	"api-contract-go/cmd/api_contract"
	"api-contract-go/cmd/api_contract/test/helpers"
  "github.com/stretchr/testify/assert"
)

func init() {
  api_contract_test_helpers.ChangeDirectoryToProjectRoot()
}

func TestExpectValidPayload(t *testing.T) {
  expectValidPayloadToPass(t)
  expectInvalidPayloadToThrow(t)
}

func expectInvalidPayloadToThrow(t *testing.T) {
  errorCalled := false
  _error := func (args ...any) {
    errorCalled = true
  }

  api_contract.InternalValidatePayload = func (bytes []byte, httpMethod string, endpoint string) (bool, string) {
    return true, ""
  }
  api_contract.ExpectValidPayload([]byte(""), "POST", "/api/v1/test", _error)
  assert.False(t, errorCalled)
}

func expectValidPayloadToPass(t *testing.T) {
  errorCalled := false
  _error := func (args ...any) {
    errorCalled = true
  }

  api_contract.InternalValidatePayload = func (bytes []byte, httpMethod string, endpoint string) (bool, string) {
    return false, "something went wrong"
  }
  api_contract.ExpectValidPayload([]byte(""), "POST", "/api/v1/test", _error)
  assert.True(t, errorCalled)
}

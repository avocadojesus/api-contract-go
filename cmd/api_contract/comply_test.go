package api_contract_test

import (
  "testing"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract/test/helpers"
  "github.com/stretchr/testify/assert"
)

func init() {
  test_helpers.ChangeDirectoryToProjectRoot()
}

func TestComply(t *testing.T) {
  expectValidPayloadToPass(t)
  expectInvalidPayloadToThrow(t)
}

func expectInvalidPayloadToThrow(t *testing.T) {
  errorCalled := false
  _error := func (args ...any) {
    errorCalled = true
  }

  api_contract.InternalValidate = func (bytes []byte, httpMethod string, endpoint string) (bool, string) {
    return true, ""
  }
  api_contract.Comply([]byte(""), "POST", "/api/v1/test", _error)
  assert.False(t, errorCalled)
}

func expectValidPayloadToPass(t *testing.T) {
  errorCalled := false
  _error := func (args ...any) {
    errorCalled = true
  }

  api_contract.InternalValidate = func (bytes []byte, httpMethod string, endpoint string) (bool, string) {
    return false, "something went wrong"
  }
  api_contract.Comply([]byte(""), "POST", "/api/v1/test", _error)
  assert.True(t, errorCalled)
}

package api_contract_test

import (
  "testing"
  "fmt"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract/test/helpers"
)

func TestValidate(t *testing.T) {
  expectValidPayload(t, "nesting/basic")
  expectInvalidPayload(t, "nesting/basic/invalid/extra_params")
  expectInvalidPayload(t, "nesting/basic/invalid/nested_extra_params")

  expectValidPayload(t, "nesting/single_nested")
  expectInvalidPayload(t, "nesting/single_nested/invalid/extra_params")
  expectInvalidPayload(t, "nesting/single_nested/invalid/nested_extra_params")
  expectInvalidPayload(t, "nesting/single_nested/invalid/bad_type_on_nested_param")

  expectValidPayload(t, "nesting/double_nested")
  expectInvalidPayload(t, "nesting/double_nested/invalid/extra_params")
  expectInvalidPayload(t, "nesting/double_nested/invalid/nested_extra_params")
  expectInvalidPayload(t, "nesting/double_nested/invalid/bad_type_on_nested_param")

  expectValidPayload(t, "datatypes/bool")
  expectInvalidPayload(t, "datatypes/bool/invalid/bad_data_type")
  expectInvalidPayload(t, "datatypes/bool/invalid/bad_array_data_type")

  expectValidPayload(t, "datatypes/date")
  expectValidPayload(t, "datatypes/date/mmddyyyy")
  expectValidPayload(t, "datatypes/date/mmddyy")
  expectValidPayload(t, "datatypes/date/yyyymmdd")
  expectValidPayload(t, "datatypes/date/yymmdd")
  expectInvalidPayload(t, "datatypes/date/invalid/bad_data_type")
  expectInvalidPayload(t, "datatypes/date/invalid/bad_array_data_type")

  expectValidPayload(t, "datatypes/datetime")
  expectValidPayload(t, "datatypes/datetime/ansic")
  expectValidPayload(t, "datatypes/datetime/iso861")
  expectValidPayload(t, "datatypes/datetime/unix_date")
  expectValidPayload(t, "datatypes/datetime/ruby_date")
  expectValidPayload(t, "datatypes/datetime/rfc822")
  expectValidPayload(t, "datatypes/datetime/rfc822z")
  expectValidPayload(t, "datatypes/datetime/rfc850")
  expectValidPayload(t, "datatypes/datetime/rfc1123")
  expectValidPayload(t, "datatypes/datetime/rfc1123z")
  expectValidPayload(t, "datatypes/datetime/rfc3339")
  expectValidPayload(t, "datatypes/datetime/rfc3339_nano")
  expectValidPayload(t, "datatypes/datetime/kitchen")
  expectValidPayload(t, "datatypes/datetime/stamp")
  expectValidPayload(t, "datatypes/datetime/stamp_milli")
  expectValidPayload(t, "datatypes/datetime/stamp_micro")
  expectValidPayload(t, "datatypes/datetime/stamp_nano")
  expectInvalidPayload(t, "datatypes/datetime/invalid/bad_data_type")
  expectInvalidPayload(t, "datatypes/datetime/invalid/bad_array_data_type")

  expectValidPayload(t, "datatypes/number")
  expectInvalidPayload(t, "datatypes/number/invalid/bad_data_type")
  expectInvalidPayload(t, "datatypes/number/invalid/bad_array_data_type")

  expectValidPayload(t, "datatypes/string")
  expectValidPayload(t, "datatypes/string/email")
  expectValidPayload(t, "datatypes/string/name")
  expectInvalidPayload(t, "datatypes/string/invalid/bad_data_type")
  expectInvalidPayload(t, "datatypes/string/invalid/bad_array_data_type")
}

func expectValidPayload(t *testing.T, endpointStubFolder string) {
  text := api_contract.ReadJSON(fmt.Sprintf("./cmd/api_contract/test/endpoint_stubs/%s/response.json", endpointStubFolder))

  test_helpers.MockJSONRead(fmt.Sprintf("./cmd/api_contract/test/endpoint_stubs/%s/endpoints.json", endpointStubFolder))
  passedValidation, reason := api_contract.Validate([]byte(text), "POST", "/api/v1/test")
  test_helpers.RestoreJSONRead()

  if !passedValidation {
    t.Error(fmt.Sprintf("Failed to validate %s, reason: ", endpointStubFolder), reason)
  }
}

func expectInvalidPayload(t *testing.T, endpointStubFolder string) {
  text := api_contract.ReadJSON(fmt.Sprintf("./cmd/api_contract/test/endpoint_stubs/%s/response.json", endpointStubFolder))

  test_helpers.MockJSONRead(fmt.Sprintf("./cmd/api_contract/test/endpoint_stubs/%s/endpoints.json", endpointStubFolder))
  passedValidation, _ := api_contract.Validate([]byte(text), "POST", "/api/v1/test")
  test_helpers.RestoreJSONRead()

  if passedValidation {
    t.Error(fmt.Sprintf("Expected %s to fail, but it didn't", endpointStubFolder))
  }
}


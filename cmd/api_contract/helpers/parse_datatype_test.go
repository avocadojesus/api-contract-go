package helpers_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/avocadojesus/api-contract-go/cmd/api_contract/helpers"
)

func TestParseDatatype(t *testing.T) {
  assertDatatypeMatch(t, "datetime", "datetime", []string{}, false, false)
  assertDatatypeMatch(t, "datetime[]", "datetime", []string{}, true, false)
  assertDatatypeMatch(t, "datetime:ansic[]", "datetime", []string{"ansic"}, true, false)
  assertDatatypeMatch(t, "datetime:ansic:another[]", "datetime", []string{"ansic", "another"}, true, false)
}

func assertDatatypeMatch(t *testing.T, str string, expectedDatatype string, expectedDecorators []string, expectedIsArray bool, expectedIsOptional bool) {
  dataType, decorators, isArray, isOptional := helpers.ParseDatatype(str)
  assert.Equal(t, dataType, expectedDatatype)
  assert.Equal(t, decorators, expectedDecorators)
  assert.Equal(t, isArray, expectedIsArray)
  assert.Equal(t, isOptional, expectedIsOptional)
}

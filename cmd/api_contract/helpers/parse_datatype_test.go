package helpers_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/avocadojesus/api-contract-go/cmd/api_contract/helpers"
)

// supports:
// string
// string[]
// bool
// bool[]
// number
// number[]
// date
// date[]
// datetime
// datetime[]

// special datetime decorators
//
// datatype decorators are able to be used with array types as well,
// and the chains will be read prior to the array brackets, and are to
// be defined in snake_case, like so:
//
// datetime:unix_date[]
//
// supports:
// datetime:ansic
// datetime:unix_date
// datetime:ruby_date
// datetime:rfc822
// datetime:rfc822z
// datetime:rfc850
// datetime:rfc1123
// datetime:rfc3339
// datetime:rfc3339nano
// datetime:kitchen
// datetime:stamp
// datetime:stamp_milli
// datetime:stamp_micro
// datetime:stamp_nano

func TestParseDatatype(t *testing.T) {
  assertDatatypeMatch(t, "datetime", "datetime", []string{}, false)
  assertDatatypeMatch(t, "datetime[]", "datetime", []string{}, true)
  assertDatatypeMatch(t, "datetime:ansic[]", "datetime", []string{"ansic"}, true)
  assertDatatypeMatch(t, "datetime:ansic:another[]", "datetime", []string{"ansic", "another"}, true)
}

func assertDatatypeMatch(t *testing.T, str string, expectedDatatype string, expectedDecorators []string, expectedIsArray bool) {
  dataType, decorators, isArray := helpers.ParseDatatype(str)
  assert.Equal(t, dataType, expectedDatatype)
  assert.Equal(t, decorators, expectedDecorators)
  assert.Equal(t, isArray, expectedIsArray)
}

package helpers

import (
  "strings"
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

func ParseDatatype(datatype string) (string, []string, bool) {
  dataTypeWithoutArray := strings.ReplaceAll(datatype, "[]", "")
  parts := strings.Split(dataTypeWithoutArray, ":")
  perceivedDatatype := strings.ReplaceAll(parts[0], "[]", "")
  decorators := parts[1:]
  isArray := strings.Contains(datatype, "[]")

  return perceivedDatatype, decorators, isArray
}

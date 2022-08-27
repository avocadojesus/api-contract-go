package helpers

import (
  "strings"
)

func ParseDatatype(datatype string) (string, []string, bool, bool) {
  dataTypeWithoutArray := strings.ReplaceAll(datatype, "[]", "")
  parts := strings.Split(dataTypeWithoutArray, ":")
  perceivedDatatype := strings.ReplaceAll(parts[0], "[]", "")
  decorators := parts[1:]
  isArray := strings.Contains(datatype, "[]")
  isOptional := SliceContains(decorators, "optional")

  return perceivedDatatype, decorators, isArray, isOptional
}

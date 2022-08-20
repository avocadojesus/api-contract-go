package helpers

import (
  "fmt"
  "strings"
)

func IsMap(x interface{}) bool {
  t := fmt.Sprintf("%T", x)
  return strings.HasPrefix(t, "map[")
}

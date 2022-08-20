package api_contract_helpers

import (
  "fmt"
  "io/ioutil"
  "log"
)

func ReadJSON(path string) []byte {
  d, err := ioutil.ReadFile(path)

  if err != nil {
    log.Fatal(fmt.Sprintf("Missing JSON file: %s", path), err)
  }

  return d
}


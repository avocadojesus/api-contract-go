package api_contract_helpers

import (
  "fmt"
  "io/ioutil"
  "log"
  "path/filepath"
  "os"
)

func ReadJSON(path string) []byte {
  f, _ := os.Getwd()
  fmt.Println(filepath.Base(f))
  fmt.Println(filepath.Dir(f))

  d, err := ioutil.ReadFile(path)

  if err != nil {
    log.Fatal(fmt.Sprintf("Missing JSON file: %s", path), err)
  }

  return d
}


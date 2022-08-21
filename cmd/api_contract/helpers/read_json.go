package helpers

import (
  "fmt"
  "io/ioutil"
  "log"
)

func ReadJSON(path string) []byte {
  projectRoot := FindProjectRoot()
  fullPath := projectRoot + "/" + path
  d, err := ioutil.ReadFile(fullPath)

  if err != nil {
    log.Fatal(fmt.Sprintf("Missing JSON filezzzzz: %s\n", projectRoot), err)
  }

  return d
}

package test_helpers

import (
  "io/ioutil"
  "log"
  "fmt"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract/helpers"
)

var _readJSON func(_path string) []byte
func MockJSONRead(path string) {
  _readJSON = api_contract.ReadJSON

  api_contract.ReadJSON = func(_path string) []byte {
    projectRoot := helpers.FindProjectRoot()
    fullPath := projectRoot + "/" + path
    d, err := ioutil.ReadFile(fullPath)

    if err != nil {
      log.Fatal(fmt.Sprintf("Missing JSON file (from MockJSONRead): %s", fullPath), err)
    }

    return d
  }
}

func RestoreJSONRead() {
  api_contract.ReadJSON = _readJSON
}

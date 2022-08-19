package api_contract_test_helpers

import (
  "io/ioutil"
  "log"
  "fmt"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract"
)

var _readJSON func(_path string) []byte
func MockJSONRead(path string) {
  _readJSON = api_contract.ReadJSON

  api_contract.ReadJSON = func(_path string) []byte {
    d, err := ioutil.ReadFile(path)

    if err != nil {
      log.Fatal(fmt.Sprintf("Missing JSON file: %s", path), err)
    }

    return d
  }
}

func RestoreJSONRead() {
  api_contract.ReadJSON = _readJSON
}

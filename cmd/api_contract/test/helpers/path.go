package test_helpers

import (
  "runtime"
  "path"
  "os"
)

func ChangeDirectoryToProjectRoot() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

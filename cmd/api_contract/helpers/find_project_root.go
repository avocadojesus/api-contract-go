package helpers

import (
  "os"
  "errors"
)

func FindProjectRoot() string {
  startPath, _ := os.Getwd()
  if fileExists(startPath + "/go.mod") {
    return startPath
  }
  return findProjectRootRecursive(startPath + "/..")
}

func findProjectRootRecursive(path string) string {
  if fileExists(path + "/go.mod") {
    return path
  }
  return findProjectRootRecursive(path + "/..")
}

func fileExists(path string) bool {
  if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
    return false
  }
  return true
}

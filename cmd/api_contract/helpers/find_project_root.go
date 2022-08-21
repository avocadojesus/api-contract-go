package helpers

import (
  "os"
  "errors"
  "regexp"
)

func FindProjectRoot() string {
  startPath, _ := os.Getwd()
  if fileExists(startPath + "/go.mod") {
    return startPath
  }
  return findProjectRootRecursive(upOnDirectory(startPath))
}

func findProjectRootRecursive(path string) string {
  if fileExists(path + "/go.mod") {
    return path
  }

  if (path == upOnDirectory(path)) {
    panic("ERROR (api-contract-go): FindProjectRoot failed because it recursively searches your filesystem for a go.mod file.\n" +
      "If there is no go.mod file in your project root, this function will fail, though it never should, since \n" +
      "typically a go.mod file should be present in order to bring in a package in the first place.")
  }

  return findProjectRootRecursive(upOnDirectory(path))
}

func fileExists(path string) bool {
  if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
    return false
  }
  return true
}

func upOnDirectory(path string) string {
  re := regexp.MustCompile(`\/[^/]*$`)
  return re.ReplaceAllString(path, "")
}

package api_contract_helpers

func SliceContains(arr []string, matchString string) bool {
  for _, a := range arr {
    if a == matchString {
      return true
    }
  }
  return false
}

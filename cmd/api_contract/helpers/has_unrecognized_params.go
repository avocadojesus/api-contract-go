package helpers

func HasUnrecognizedParams(results map[string]interface{}, endpoints map[string]interface{}) bool {
  for key, val := range results {
    if IsMap(val) && IsMap(endpoints[key]) {
      return HasUnrecognizedParams(results[key].(map[string]interface{}), endpoints[key].(map[string]interface{}))
    } else {
      if results[key] != nil && endpoints[key] == nil {
        return false
      }
    }
  }
  return true
}

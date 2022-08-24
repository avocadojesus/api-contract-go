package validators

func ValidateBool(
  param string,
  paramType interface{},
  results map[string]interface{},
  typeOfReturnedValue string,
  isArray bool,
) bool {
  if isArray {
    return validateBoolArray(results[param].([]interface{}))
  } else {
    return typeOfReturnedValue == paramType
  }
}

func validateBoolArray(arr []interface{}) bool {
  return CheckArrayForType(arr, "bool")
}

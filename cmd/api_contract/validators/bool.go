package validators

func ValidateBoolArray(arr []interface{}) bool {
  return CheckArrayForType(arr, "bool")
}

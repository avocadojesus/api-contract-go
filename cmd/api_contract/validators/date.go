package validators

import (
  "time"
  "fmt"
  "regexp"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract/helpers"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract/config"
)

func ValidateDate(
  param string,
  paramType interface{},
  results map[string]interface{},
  decorators []string,
  typeOfReturnedValue string,
  isArray bool,
  conf config.ApiContractConfig,
) bool {
  if (len(decorators) == 0) {
    if isArray {
      return validateDateArray(results[param].([]interface{}))
    } else {
      return validateDateBasic(results[param].(interface{}))
    }
  }

  format := findDateFormat(decorators)
  if format == "" {
    return false
  }

  if isArray {
    return validateDateArrayCustomFormat(results[param].([]interface{}), format)
  } else {
    return validateDateCustomFormat(results[param].(interface{}), format)
  }
}

func validateDateBasic(date interface{}) bool {
  _, err := time.Parse(DATE_FORMAT, fmt.Sprintf("%s", date))
  return err == nil
}

func validateDateCustomFormat(date interface{}, format string) bool {
  matchFound, err := regexp.MatchString(format, fmt.Sprintf("%s", date))
  if err != nil {
    return false
  }
  return matchFound
}

func validateDateArray(arr []interface{}) bool {
  for _, item := range arr {
    if !validateDateBasic(item) {
      return false
    }
  }
  return true
}

func validateDateArrayCustomFormat(arr []interface{}, format string) bool {
  for _, item := range arr {
    if !validateDateCustomFormat(item, format) {
      return false
    }
  }
  return true
}

func findDateFormat(arr []string) string {
  if SliceContains(arr, "mmddyyyy") || SliceContains(arr, "MMDDYYYY") {
    return helpers.Date.MMDDYYYY
  } else if SliceContains(arr, "mmddyy") || SliceContains(arr, "MMDDYY") {
    return helpers.Date.MMDDYY
  } else if SliceContains(arr, "yyyymmdd") || SliceContains(arr, "YYYYMMDD") {
    return helpers.Date.YYYYMMDD
  } else if SliceContains(arr, "yymmdd") || SliceContains(arr, "YYMMDD") {
    return helpers.Date.YYMMDD
  } else {
    return ""
  }
}


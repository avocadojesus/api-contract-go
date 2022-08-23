package validators

import (
  "time"
  "fmt"
  "regexp"
	"github.com/avocadojesus/api-contract-go/cmd/api_contract/helpers"
)

func ValidateDate(date interface{}) bool {
  _, err := time.Parse(DATE_FORMAT, fmt.Sprintf("%s", date))
  return err == nil
}

func ValidateDateCustomFormat(date interface{}, format string) bool {
  matchFound, err := regexp.MatchString(format, fmt.Sprintf("%s", date))
  if err != nil {
    return false
  }
  return matchFound
}

func ValidateDateArray(arr []interface{}) bool {
  for _, item := range arr {
    if !ValidateDate(item) {
      return false
    }
  }
  return true
}

func ValidateDateArrayCustomFormat(arr []interface{}, format string) bool {
  for _, item := range arr {
    if !ValidateDateCustomFormat(item, format) {
      return false
    }
  }
  return true
}

func FindDateFormat(arr []string) string {
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


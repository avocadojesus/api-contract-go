package validators

import (
  "time"
  "fmt"
)

func ValidateDatetime(date interface{}) bool {
  _, err := time.Parse(DATETIME_FORMAT, fmt.Sprintf("%s", date))
  return err == nil
}

func ValidateDatetimeCustomFormat(date interface{}, format string) bool {
  _, err := time.Parse(format, fmt.Sprintf("%s", date))
  return err == nil
}

func ValidateDatetimeArray(arr []interface{}) bool {
  for _, item := range arr {
    if !ValidateDatetime(item) {
      return false
    }
  }
  return true
}

func ValidateDatetimeArrayCustomFormat(arr []interface{}, format string) bool {
  for _, item := range arr {
    if !ValidateDatetimeCustomFormat(item, format) {
      return false
    }
  }
  return true
}

func FindDatetimeFormat(arr []string) string {
  if SliceContains(arr, "ansic") {
    return time.ANSIC
  } else if (SliceContains(arr, "iso861") || SliceContains(arr, "ISO861")) {
    return time.RFC3339
  } else if (SliceContains(arr, "unix_date") || SliceContains(arr, "unix")) {
    return time.UnixDate
  } else if (SliceContains(arr, "ruby_date") || SliceContains(arr, "ruby")) {
    return time.RubyDate
  } else if (SliceContains(arr, "rfc822") || SliceContains(arr, "RFC822")) {
    return time.RFC822
  } else if (SliceContains(arr, "rfc822z") || SliceContains(arr, "RFC822Z")) {
    return time.RFC822Z
  } else if (SliceContains(arr, "rfc850") || SliceContains(arr, "RFC850")) {
    return time.RFC850
  } else if (SliceContains(arr, "rfc1123") || SliceContains(arr, "RFC1123")) {
    return time.RFC1123
  } else if (SliceContains(arr, "rfc1123z") || SliceContains(arr, "RFC1123Z")) {
    return time.RFC1123Z
  } else if (SliceContains(arr, "rfc3339") || SliceContains(arr, "RFC3339")) {
    return time.RFC3339
  } else if (SliceContains(arr, "rfc3339_nano") || SliceContains(arr, "RFC3339Nano")) {
    return time.RFC3339Nano
  } else if (SliceContains(arr, "kitchen") || SliceContains(arr, "Kitchen")) {
    return time.Kitchen
  } else if
    SliceContains(arr, "stamp") ||
      SliceContains(arr, "Stamp") ||
      SliceContains(arr, "timestamp") ||
      SliceContains(arr, "Timestamp") {
    return time.Stamp
  } else if
    SliceContains(arr, "stamp_milli") ||
      SliceContains(arr, "StampMilli") ||
      SliceContains(arr, "timestamp_milli") ||
      SliceContains(arr, "TimestampMilli") {
    return time.StampMilli
  } else if
    SliceContains(arr, "stamp_micro") ||
      SliceContains(arr, "StampMicro") ||
      SliceContains(arr, "timestamp_micro") ||
      SliceContains(arr, "TimestampMicro") {
    return time.StampMicro
  } else if
    SliceContains(arr, "stamp_nano") ||
      SliceContains(arr, "StampNano") ||
      SliceContains(arr, "timestamp_nano") ||
      SliceContains(arr, "TimestampNano") {
    return time.StampNano
  } else {
    return ""
  }
}


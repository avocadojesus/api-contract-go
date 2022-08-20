package helpers

type IDate struct {
  MMDDYYYY string
  MMDDYY string
  YYYYMMDD string
  YYMMDD string
}

var Date = IDate{
  MMDDYYYY: `\d\d[-/]?\d\d[-/]?\d\d\d\d`,
  MMDDYY: `\d\d[-/]?\d\d[-/]?\d\d`,
  YYYYMMDD: `\d\d\d\d[-/]?\d\d[-/]?\d\d`,
  YYMMDD: `\d\d[-/]?\d\d[-/]?\d\d`,
}

package parser

import "time"

type timeString struct {
	Date                  time.Time
	timeOnlyFormat        string
	withMilliSecondFormat string
	withDateFormat        string
}

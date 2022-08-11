package parser

import (
	"fmt"
	"strconv"
	"time"
)

func TimeStringDateOnly(dur time.Time) (format string) {
	trs := newTimeString(dur).WithDate(" ", " ").Output()
	return trs[:len(trs)-1]
}

func TimeStringTimeOnly(dur time.Time) (format string) {
	return newTimeString(dur).WithTime().Output()
}

func TimeStringTimeOnlyWithMilliSecond(s time.Time) (ts string) {
	return newTimeString(s).WithTime().WithMilliSecond().Output()
}

func TimeStringWithDate(dur time.Time) (format string) {
	return newTimeString(dur).WithTime().WithDate(" ", " ").Output()
}

func TimeStringWitLongDate(dur time.Time) (format string) {
	return newTimeString(dur).WithTime().WithLongDate(" ", " ").Output()
}

func (cc *timeString) WithLongDate(sp, sp2 string) *timeString {
	cc.withDateFormat = fmt.Sprintf("02%sJanuary%s2006%s", sp, sp, sp2)
	return cc
}

func (cc *timeString) WithDate(sp, sp2 string) *timeString {
	cc.withDateFormat = fmt.Sprintf("02%sJan%s2006%s", sp, sp, sp2)
	return cc
}

func (cc *timeString) WithTime() *timeString {
	cc.timeOnlyFormat = "15:04:05"
	return cc
}

func (cc *timeString) WithMilliSecond() *timeString {
	cc.withMilliSecondFormat = strconv.Itoa(cc.Date.Nanosecond() / 1000000)
	if len(cc.withMilliSecondFormat) < 3 {
		rem := 3 - len(cc.withMilliSecondFormat)
		for i := 0; i < rem; i++ {
			cc.withMilliSecondFormat += `0`
		}
	}
	cc.withMilliSecondFormat = "." + cc.withMilliSecondFormat

	return cc
}

func (cc *timeString) Output() (sr string) {
	return fmt.Sprintf("%s%s",
		cc.Date.Format(fmt.Sprintf("%s%s",
			cc.withDateFormat, cc.timeOnlyFormat)),
		cc.withMilliSecondFormat)
}

func newTimeString(tt time.Time) *timeString {
	return &timeString{Date: tt}
}

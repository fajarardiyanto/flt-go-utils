package daterange

import (
	"log"
	"time"
)

type DatesRange struct {
	FirstDay time.Time `json:"first_day"`
	LastDay  time.Time `json:"last_day"`
}

func DateRange(start, end int64) []DatesRange {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println(err)
		return nil
	}

	var tempFirstDate time.Time
	var tempLastDate time.Time
	var datesRanges []DatesRange

	i := 1
	for {
		if tempFirstDate.IsZero() {
			if end > 0 {
				now := time.Unix(end, 0).In(location)
				tempFirstDate = FirstDayDate(now)
			} else {
				now := time.Now().In(location)
				tempFirstDate = FirstDayDate(now)
			}
		} else {
			tempFirstDate = tempFirstDate.Add(-(24 * time.Hour))
		}

		tempLastDate = LastDayDate(tempFirstDate)

		datesRanges = append(datesRanges, DatesRange{tempFirstDate, tempLastDate})

		if start > 0 && end > 0 {
			if start >= tempFirstDate.Unix() {
				break
			}
		} else {
			if i >= 7 {
				break
			}
		}

		i++
	}

	return datesRanges
}

func FirstDayDate(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func LastDayDate(t time.Time) time.Time {
	year, month, day := t.Date()
	firstDay := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return firstDay.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
}

package times

import (
	"fmt"
	"time"
)

//go:generate stringer -type Repeat
type Repeat int

const (
	Never Repeat = iota
	Daily
	Weekly
	BiWeekly
	Monthly
	Quarterly
	Yearly
)

func (r Repeat) IsValid() bool {
	switch r {
	case Never, Daily, Weekly, Monthly, Yearly:
		return true
	default:
		return false
	}
}

func RepeatSchedule(repeat Repeat, start time.Time, before time.Time) []time.Time {
	var res []time.Time
	t := start
	for t.Before(before) {
		res = append(res, t)
		if repeat == Never {
			return res
		}
		t = RepeatNext(repeat, t)
	}
	return res
}

func RepeatNext(repeat Repeat, start time.Time) time.Time {
	switch repeat {
	case Daily:
		return start.AddDate(0, 0, 1)
	case Weekly:
		return start.AddDate(0, 0, 7)
	case BiWeekly:
		return start.AddDate(0, 0, 14)
	case Monthly:
		return start.AddDate(0, 1, 0)
	case Quarterly:
		return start.AddDate(0, 3, 0)
	case Yearly:
		return start.AddDate(1, 0, 0)
	default:
		panic(fmt.Sprintf("invalid repeat %d: %v", repeat, repeat))
	}
}

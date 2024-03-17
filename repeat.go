package times

import "time"

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

func RepeatSchedule(repeat Repeat, initialStart time.Time, before time.Time) []time.Time {
	var res []time.Time
	t := initialStart
	for t.Before(before) {
		res = append(res, t)
		switch repeat {
		case Daily:
			t = t.AddDate(0, 0, 1)
		case Weekly:
			t = t.AddDate(0, 0, 7)
		case Monthly:
			t = t.AddDate(0, 1, 0)
		case Quarterly:
			t = t.AddDate(0, 3, 0)
		case Yearly:
			t = t.AddDate(1, 0, 0)
		default:
			return res
		}
	}
	return res
}

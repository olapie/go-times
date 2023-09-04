package times

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

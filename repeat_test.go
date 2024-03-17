package times

import (
	"testing"
	"time"
)

func TestRepeatSchedule(t *testing.T) {
	start := time.Date(2025, time.January, 2, 0, 0, 0, 0, time.Local)
	t.Log(start, start.Weekday())
	t.Log(RepeatSchedule(Weekly, start, time.Now().AddDate(2, 0, 0)))
}

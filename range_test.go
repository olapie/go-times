package timex_test

import (
	"testing"
	"time"

	"code.olapie.com/log"
	"code.olapie.com/timex"
)

func TestRange_SplitInDay(t *testing.T) {
	tz := time.FixedZone("PST", -7*3600)
	begin := time.Date(2002, 5, 3, 17, 0, 0, 0, tz)
	end := time.Date(2002, 5, 3, 18, 0, 0, 0, tz)
	r := timex.NewRange(begin, end)
	for _, dr := range r.SplitInDay() {
		log.S().Debug(dr.Begin(), dr.End())
	}
}

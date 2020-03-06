package timezone

import (
	"testing"
)

func TestTimezone(t *testing.T) {
	start, err := DayDateToUnix(9,
		"2020-01-17", 0)
	t.Logf("time: %d err:%+v", start, err)

	day, hour, err := UnixToDayDate(start, 16)
	t.Logf("day: %s hour:%d err:%+v", day, hour, err)

	day, hour, err = UnixToDayDate(start, 24)
	t.Logf("day: %s hour:%d err:%+v", day, hour, err)

	day, hour, err = UnixToDayDate(start, 8)
	t.Logf("day: %s hour:%d err:%+v", day, hour, err)
}

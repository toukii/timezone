package timezone

/*
https://github.com/tkuchiki/go-timezone/blob/master/timezone.go
http://www.shijian.cc/shiqu/

东西十二区 ( UTC +12 )
东十一区 ( UTC +11 )
东十区 ( UTC +10 )
东九区 ( UTC +9 )
东八区 ( UTC +8 )
东七区 ( UTC +7 )
东六区 ( UTC +6 )
东五区 ( UTC +5 )
东四区 ( UTC +4 )
东三区 ( UTC +3 )
东二区 ( UTC +2 )
东一区 ( UTC +1 )
中时区 ( UTC +0 )
西一区 ( UTC -1 )
西二区 ( UTC -2 )
西三区 ( UTC -3 )
西四区 ( UTC -4 )
西五区 ( UTC -5 )
西六区 ( UTC -6 )
西七区 ( UTC -7 )
西八区 ( UTC -8 )
西九区 ( UTC -9 )
西十区 ( UTC -10 )
西十一区 ( UTC -11 )
*/

import (
	"fmt"
	"time"
)

var (
	Zonenames = map[int32]string{
		24: "中时区",
		1:  "东一区",
		2:  "东二区",
		3:  "东三区",
		4:  "东四区",
		5:  "东五区",
		6:  "东六区",
		7:  "东七区",
		8:  "东八区",
		9:  "东九区",
		10: "东十区",
		11: "东十一区",
		12: "东西十二区",
		13: "西一区",
		14: "西二区",
		15: "西三区",
		16: "西四区",
		17: "西五区",
		18: "西六区",
		19: "西七区",
		20: "西八区",
		21: "西九区",
		22: "西十区",
		23: "西十一区",
	}

	timezones = map[int32]string{
		24: "+00:00",
		1:  "+01:00",
		2:  "+02:00",
		3:  "+03:00",
		4:  "+04:00",
		5:  "+05:00",
		6:  "+06:00",
		7:  "+07:00",
		8:  "+08:00",
		9:  "+09:00",
		10: "+10:00",
		11: "+11:00",
		12: "+12:00",
		13: "-01:00",
		14: "-02:00",
		15: "-03:00",
		16: "-04:00",
		17: "-05:00",
		18: "-06:00",
		19: "-07:00",
		20: "-08:00",
		21: "-09:00",
		22: "-10:00",
		23: "-11:00",
	}

	zones = map[int32]string{
		// 0: "Europe/London",
		// 8: "Asia/Shanghai",
		// 9: "Asia/Tokyo",

		24: "Etc/GMT+0",
		1:  "Etc/GMT-1",
		2:  "Etc/GMT-2",
		3:  "Etc/GMT-3",
		4:  "Etc/GMT-4",
		5:  "Etc/GMT-5",
		6:  "Etc/GMT-6",
		7:  "Etc/GMT-7",
		8:  "Etc/GMT-8",
		9:  "Etc/GMT-9",
		10: "Etc/GMT-10",
		11: "Etc/GMT-11",
		12: "Etc/GMT-12",
		13: "Etc/GMT+1",
		14: "Etc/GMT+2",
		15: "Etc/GMT+3",
		16: "Etc/GMT+4",
		17: "Etc/GMT+5",
		18: "Etc/GMT+6",
		19: "Etc/GMT+7",
		20: "Etc/GMT+8",
		21: "Etc/GMT+9",
		22: "Etc/GMT+10",
		23: "Etc/GMT+11",
	}
)

// time.Parse("2006-01-02T15:04:05Z07:00", "2018-10-01T16:27:00+08:00")
func ParseUTC(date string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05Z07:00", date)
}

func DayDateToUnix(timezone int32, dayDate string, hour int) (int64, error) {
	if timezone <= 0 || timezone > 24 {
		return 0, fmt.Errorf("timezone not exist.")
	}

	datestr := fmt.Sprintf("%sT%d:00:00%s", dayDate, hour, timezones[timezone])
	date, err := ParseUTC(datestr)
	if err != nil {
		return 0, err
	}

	fmt.Printf("%s %s %d ==> %d\n", Zonenames[timezone], dayDate, hour, date.Unix())

	return date.Unix(), nil
}

func UnixToDayDate(unix int64, timezone int32) (dayDate string, hour int32, err error) {
	if timezone <= 0 || timezone > 24 {
		return "", 0, fmt.Errorf("timezone not exist.")
	}
	defer func() {
		fmt.Printf("%d ==> %s %s %d\n", unix, Zonenames[timezone], dayDate, hour)
	}()

	zone := zones[timezone]
	loc, err := time.LoadLocation(zone)
	if err != nil {
		return "", 0, err
	}
	loctime := time.Unix(unix, 0).In(loc)
	return loctime.Format("2006-01-02"), int32(loctime.Hour()), nil
}

package utils

import (
	"fmt"
	"time"
)

var DbTimeFormat = "2006-01-02 15:04:05"

// 2006-01-02 15:04:05 -> time
func ParseWithLocal(timeStr string) (time.Time, error) {
	locationName := "Local"
	if l, err := time.LoadLocation(locationName); err != nil {
		println(err.Error())
		return time.Time{}, err
	} else {
		lt, _ := time.ParseInLocation(DbTimeFormat, timeStr, l)
		fmt.Println(locationName, lt)
		return lt, nil
	}
}

// yyyyMMDD
func GetCurDate() string {
	t := time.Now()
	return t.Format("20060102")

}

func SecondsUntilEndOfToday() int {
	// 获取当前时间
	currentTime := time.Now()
	// 获取今晚23:59:59的时间
	endOfDay := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())
	// 计算相差的时间
	duration := endOfDay.Sub(currentTime)
	// 返回相差的秒数
	return int(duration.Seconds())
}

// eg: return:1552893276000
func GenTimestampMs(dob string) int64 {
	return GenTimestamp(dob) * 1000
}

// eg: return:1552893276
func GenTimestamp(dob string) int64 {
	t, err := ParseWithLocal(dob)
	if err != nil {
		//fmt.Println(err)
		return 0
	}
	return t.Unix()
}

// 2006-01-02 15:04:05 转化为时间戳
func GenDbTime(ctime time.Time) string {
	return ctime.Format(DbTimeFormat)
}

// 1552893276000 转化为  2006-01-02 15:04:05
func GenTimeFromTsMs(timestamp int64) string {
	loc, _ := time.LoadLocation("Asia/Shanghai") // 设置时区为东八区
	t := time.Unix(timestamp/1000, 0).In(loc)    // 转换为time.Time类型并应用时区
	return t.Format(DbTimeFormat)
}

// 2006-01-02 03:04:05
func CalculateAge(dob string) int32 {
	t, err := ParseWithLocal(dob)
	if err != nil {
		//fmt.Println(err)
		return 0
	}
	now := time.Now()
	age := now.Year() - t.Year()
	if now.Month() < t.Month() || (now.Month() == t.Month() && now.Day() < t.Day()) {
		age--
	}
	return int32(age)
}

// 计算星座
func CalculateConstellation(dob string) int32 {
	t, err := ParseWithLocal(dob)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	month := t.Month()
	day := t.Day()
	switch {
	case month == time.January && day >= 20 || month == time.February && day <= 18:
		return 11 //"水瓶座" // 11
	case month == time.February && day >= 19 || month == time.March && day <= 20:
		return 12 //"双鱼座" // 12
	case month == time.March && day >= 21 || month == time.April && day <= 19:
		return 1 //"白羊座" // 1
	case month == time.April && day >= 20 || month == time.May && day <= 20:
		return 2 //"金牛座"	//
	case month == time.May && day >= 21 || month == time.June && day <= 21:
		return 3 //"双子座"
	case month == time.June && day >= 22 || month == time.July && day <= 22:
		return 4 //"巨蟹座"
	case month == time.July && day >= 23 || month == time.August && day <= 22:
		return 5 //"狮子座"
	case month == time.August && day >= 23 || month == time.September && day <= 22:
		return 6 //"处女座"
	case month == time.September && day >= 23 || month == time.October && day <= 23:
		return 7 //"天秤座"
	case month == time.October && day >= 24 || month == time.November && day <= 21:
		return 8 //"天蝎座"
	case month == time.November && day >= 22 || month == time.December && day <= 21:
		return 9 // "射手座"
	case month == time.December && day >= 22 || month == time.January && day <= 19:
		return 10 //"摩羯座"
	default:
		return 0
	}
}

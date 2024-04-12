package timeutil

import (
	"asense-comutil/comutils/characterutil"
	"fmt"
	"time"
)

// TimeFormat 将time.Time时间格式化为时间字符串"YYYY-MM-DD HH:mm:ss"
func TimeFormat(t *time.Time) string {
	if t == nil {
		return "——"
	}
	return t.Format(time.DateTime)
}

// TimeMilliFormat 将int64毫秒时间格式化为时间字符串"YYYY-MM-DD HH:mm:ss"
func TimeMilliFormat(t *int64) string {
	if t == nil || *t == 0 {
		return "——"
	}
	return time.Unix(*t/1e3, 0).Format(time.DateTime)
}

// TimeMilliFormatToYMD 将int64毫秒时间格式化为时间字符串"YYYY-MM-DD"
func TimeMilliFormatToYMD(t *int64) string {
	if t == nil || *t == 0 {
		return "——"
	}
	return time.Unix(*t/1e3, 0).Format(time.DateOnly)
}

// SetTimeToDayEnd 将给定的时间戳将时分秒设置为23:59:59
// t: 时间戳
//
// return: 设置为当天结束的时间戳
func SetTimeToDayEnd(t int64) int64 {
	_t := time.Unix(t/1e3, 0)

	// Set time to the end of the day (23:59:59)
	endOfDay := time.Date(_t.Year(), _t.Month(), _t.Day(), 23, 59, 59, 0, _t.Location())
	return endOfDay.UnixMilli()
}

// TimeCode 生成时间编码
// num: 编码数字
//
// return: 时间编码
func TimeCode(num int) string {
	return characterutil.StitchingBuilderStr(time.Now().Format("20060102"), fmt.Sprintf("%08d", num))
}

// DateStampByDate 通过日期(年月日)获取时间戳
// date: 日期
//
// return: 时间戳
func DateStampByDate(date string) int64 {
	t, _ := time.ParseInLocation("2006/01/02", date, time.Local)
	return t.UnixMilli()
}

// DateSubDay 获取指定天数之前的日期
// day: 天数
//
// return: 开始日期, 结束日期
func DateSubDay(day int) (string, string) {
	return time.Now().AddDate(0, 0, -day).Format("2006/01/02"), time.Now().Format("2006/01/02")
}

// CalculateDateSubDay 日期减天(年月日时分秒)输出开始时间和结束时间的毫秒时间戳
// day: 天数
//
// return: 开始时间, 结束时间
func CalculateDateSubDay(day int) (int64, int64) {
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -day+1)

	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 999999999, endDate.Location())

	startTimestamp := startDate.UnixNano() / int64(time.Millisecond)
	endTimestamp := endDate.UnixNano() / int64(time.Millisecond)

	return startTimestamp, endTimestamp
}

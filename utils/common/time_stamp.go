package common

import "time"

// TimeToStamp 将符合格式的时间字符串转换为对应时区的时间戳
func TimeToStamp(s string) int64 {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Shanghai")                             //设置时区
	times, _ := time.ParseInLocation(timeLayout, "2022-11-30 23:59:59", loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	timeUnix := times.Unix()
	return timeUnix
}

// StampToTime 将时间戳转换为 Time
func StampToTime(i int64) time.Time {
	timeStr2 := time.Unix(i, 0)
	return timeStr2
}

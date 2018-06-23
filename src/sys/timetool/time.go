package timetool

import "time"

// GetTime ... 获取当前系统时间
func GetTime(time time.Time) string {
	return time.Format("060102150405")
}

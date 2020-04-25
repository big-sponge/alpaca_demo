package common

import "time"

var (
	BasePath       = "."
	DateTimeFormat = "2006-01-02 15:04:05"
	DateTimeFormatMST = "2006-01-02 15:04:05 -0700 MST"
)

func GetCurTime() time.Time {
	return time.Now().UTC()
}

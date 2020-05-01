package date_utils

import "time"

const (
	apiDateFormat = "02-01-2006T1:02:05Z"
	apiDbLayout   = "2006-01-02 15:02:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}
func GetNowString() string {

	return GetNow().Format(apiDateFormat)

}
func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}

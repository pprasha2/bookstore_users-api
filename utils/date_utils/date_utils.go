package date_utils

import "time"

const (
	apiDateFormat = "02-01-2006T1:02:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}
func GetNowString() string {

	return GetNow().Format(apiDateFormat)

}

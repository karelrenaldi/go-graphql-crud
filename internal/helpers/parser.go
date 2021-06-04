package helpers

import "time"

func StringToDate(value string) time.Time {
	var layoutFormat = "2006-01-02 15:04:05"
	var date, _ = time.Parse(layoutFormat, value)

	return date
}

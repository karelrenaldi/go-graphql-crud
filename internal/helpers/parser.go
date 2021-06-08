package helpers

import "time"

// StringToDate is function to parsing value with type string to type time.Time
// This function will return value with type time.Time
func StringToDate(value string) time.Time {
	var layoutFormat = "2006-01-02 15:04:05"
	var date, _ = time.Parse(layoutFormat, value)

	return date
}

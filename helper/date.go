package helper

import (
	"time"
)

// GetDateNow is used to get realtime now date
func GetDateNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GenDateNameFile function is for Generate name file base on realtime date now
func GenDateNameFile() string {
	return time.Now().Format("20060102150405")
}

// ParseShortDate Parse Very Short Date
func ParseShortDate(date string) time.Time {
	timeFormat := "2006-01-02 15:04:05"
	then, _ := time.Parse(timeFormat, date)
	return then
}

func DatetimeDiff(date1 string, date2 string, format string) float64 {
	timeFormat := "2006-01-02 15:04:05"
	currentTime := time.Now()
	loc := currentTime.Location()
	pasttime1, _ := time.ParseInLocation(timeFormat, date1, loc)
	pasttime2, _ := time.ParseInLocation(timeFormat, date2, loc)

	var diff float64

	switch format {
	case "hours":
		diff = pasttime2.Sub(pasttime1).Hours()
	case "minutes":
		diff = pasttime2.Sub(pasttime1).Minutes()
	default:
		diff = pasttime2.Sub(pasttime1).Hours()
	}

	return diff
}

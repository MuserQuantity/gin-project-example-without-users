package utils

import "time"

func GetHourTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, time.Local)
}

func GetDayTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

func GetYearTime(t time.Time) (serises []time.Time) {
	year := t.Year()
	month := t.Month()
	for i := 0; i < 12; i++ {
		month--
		if month == 0 {
			month = 12
			year--
		}
		serises = append(serises, time.Date(year, month, 0, 0, 0, 0, 0, time.Local))
	}
	return
}

package main

import (
	"time"
)

type Calendar struct {
	year, month, day int
}

func NewCalendar(thetime time.Time) (result Calendar) {
	result.year = int(thetime.Year())
	result.month = int(thetime.Month())
	result.day = int(thetime.Day())
	return result
}

func (calendar Calendar) CurrentQuarter() int {
	return (calendar.month-1)/3 + 1
}

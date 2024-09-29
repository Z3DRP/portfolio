package models

import (
	"fmt"
	"time"
)

type DailyAvailability struct {
	AvailableFrom time.Time
	AvailableTo   time.Time
}

func (d DailyAvailability) DayKey() string {
	return fmt.Sprintf(
		"%v %v %v %v",
		d.AvailableFrom.Weekday().String(),
		d.AvailableFrom.Month().String(),
		d.AvailableFrom.Day(),
		d.AvailableFrom.Year(),
	)
}

func NewTempDailyAvail() []DailyAvailability {
	return []DailyAvailability{
		{AvailableFrom: time.Date(2024, time.September, 30, 9, 0, 0, 0, time.Local), AvailableTo: time.Date(2024, time.September, 30, 15, 0, 0, 0, time.Local)},
		{AvailableFrom: time.Date(2024, time.September, 31, 9, 0, 0, 0, time.Local), AvailableTo: time.Date(2024, time.September, 31, 15, 0, 0, 0, time.Local)},
		{AvailableFrom: time.Date(2024, time.October, 1, 9, 0, 0, 0, time.Local), AvailableTo: time.Date(2024, time.October, 1, 15, 0, 0, 0, time.Local)},
		{AvailableFrom: time.Date(2024, time.October, 2, 9, 0, 0, 0, time.Local), AvailableTo: time.Date(2024, time.October, 2, 15, 0, 0, 0, time.Local)},
		{AvailableFrom: time.Date(2024, time.October, 3, 9, 0, 0, 0, time.Local), AvailableTo: time.Date(2024, time.October, 3, 15, 0, 0, 0, time.Local)},
		{AvailableFrom: nil, AvailableTo: nil},
	}
}

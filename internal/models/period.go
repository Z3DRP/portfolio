package models

import (
	"fmt"
	"time"
)

const DAY_HRS = 24

type Navigator interface {
	Next() (time.Time, time.Time)
	Prev() (time.Time, time.Time)
}

type Period struct {
	Length    int
	StartDate time.Time
	EndDate   time.Time
}

func (p Period) GetDaysInPeriod() []time.Time {
	days := make([]time.Time, 0)
	for d := 1; d <= p.Length; d++ {
		nxtDay := p.StartDate.Add(time.Duration(24*d) * time.Hour)
		if !nxtDay.After(p.EndDate) {
			days = append(days, p.StartDate)
		}
	}
	return days
}

func (p Period) FormattedDate(t time.Time) string {
	return fmt.Sprintf("%v %v %v, %v", t.Weekday(), t.Month(), t.Day(), t.Year())
}

func (p Period) Next() (time.Time, time.Time) {
	return p.StartDate.Add(time.Duration(DAY_HRS*p.Length) * time.Hour), p.EndDate.Add(time.Duration(DAY_HRS*p.Length) * time.Hour)
}

func (p Period) Prev() (time.Time, time.Time) {
	return p.StartDate.Add(time.Duration(DAY_HRS*-p.Length) * time.Hour), p.EndDate.Add(time.Duration(DAY_HRS*p.Length) * time.Hour)
}

type NextPeriod struct {
	Direction       string
	PrevPeriodStart time.Time
}

func NewPeriod(length int, start, end time.Time) Navigator {
	return &Period{
		Length:    length,
		StartDate: start,
		EndDate:   end,
	}
}

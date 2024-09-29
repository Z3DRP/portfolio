package models

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

const LOWER_TIME_BOUND = 8
const UPPER_TIME_BOUND = 17

type DailyAvailability struct {
	Day               int
	AvailableFrom     time.Time
	AvailableTo       time.Time
	AvailabilityRange *map[string]bool
}

func (d DailyAvailability) DayKey() string {
	return fmt.Sprintf(
		"%v",
		d.AvailableFrom.Weekday().String(),
	)
}

func (d DailyAvailability) StartDay() string {
	return d.AvailableFrom.Weekday().String()
}

func (d DailyAvailability) EndDay() string {
	return d.AvailableFrom.Weekday().String()
}

func (d DailyAvailability) StartTime() string {
	return fmt.Sprintf("%v:%v", d.AvailableFrom.Hour(), d.AvailableFrom.Minute())
}

func (d DailyAvailability) EndTime() string {
	return fmt.Sprintf("%v:%v", d.AvailableTo.Hour(), d.AvailableTo.Minute())
}

func (d DailyAvailability) FormattedTime() string {
	return fmt.Sprintf("%v - %v", d.StartTime(), d.EndTime())
}

func (d DailyAvailability) FormattedAvailability() string {
	return fmt.Sprintf("%v %v %v, %v - %v", d.StartDay(), d.AvailableFrom.Month().String(), d.AvailableFrom.Day(), d.AvailableFrom.Year(), d.FormattedTime())
}

func initAvailRange(from time.Time, to time.Time) (*map[string]bool, error) {
	if from.Hour() < LOWER_TIME_BOUND {
		OutOfBounds := errors.New("availabile from time out of bounds")
		return nil, OutOfBounds
	}
	if to.Hour() > UPPER_TIME_BOUND {
		OutOfBounds := errors.New("availabile to time out of bounds")
		return nil, OutOfBounds
	}
	if from.Hour() > to.Hour() {
		InvalidTimes := errors.New("from time must be before to time")
		return nil, InvalidTimes
	}

	tmRange := make(map[string]bool, UPPER_TIME_BOUND-LOWER_TIME_BOUND)
	for i := LOWER_TIME_BOUND; i <= UPPER_TIME_BOUND; i++ {
		if _, exists := tmRange[string(i)]; !exists {
			tmRange[string(i)] = false
		}
	}

	for i := to.Hour(); i <= from.Hour(); i++ {
		tmRange[string(i)] = true
	}

	return &tmRange, nil
}

type Schedule struct {
	Availability  []DailyAvailability
	DailySchedule map[string]Tasklist
}

func SortTasksByDay(s Schedule, tasks Tasklist) (map[string]Tasklist, error) {
	sortedTasks := make(map[string]Tasklist, 7)
	if len(tasks) == 0 {
		return nil, fmt.Errorf("empty task list")
	}

	for _, day := range s.Availability {
		if _, exists := s.DailySchedule[day.DayKey()]; !exists {
			s.DailySchedule[day.DayKey()] = make(Tasklist, 0)
		}
	}

	for _, task := range tasks {
		if tasksForDay, exists := s.DailySchedule[task.TaskDayKey()]; exists {
			s.DailySchedule[task.TaskDayKey()] = append(tasksForDay, task)
		}
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	keys := make([]string, 0, len(s.DailySchedule))

	for k := range sortedTasks {
		keys = append(keys, k)
	}

	for _, key := range keys {
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			// sortByTimes(&tks)
			mu.Lock()
			sort.Slice(s.DailySchedule[k], func(curIndx, nxtIndx int) bool {
				return s.DailySchedule[k][curIndx].StartTime.Compare(s.DailySchedule[k][nxtIndx].StartTime) == -1
			})
			mu.Unlock()
		}(key)
	}
	wg.Wait()

	return sortedTasks, nil
}

func NewTempDailyAvail() []DailyAvailability {
	return []DailyAvailability{
		{AvailableFrom: time.Date(2024, time.September, 30, 9, 0, 0, 0, time.Local), AvailableTo: time.Date(2024, time.September, 30, 15, 0, 0, 0, time.Local)},
		{AvailableFrom: time.Date(2024, time.September, 31, 9, 0, 0, 0, time.Local), AvailableTo: time.Date(2024, time.September, 31, 15, 0, 0, 0, time.Local)},
		{AvailableFrom: time.Date(2024, time.October, 1, 9, 0, 0, 0, time.Local), AvailableTo: time.Date(2024, time.October, 1, 15, 0, 0, 0, time.Local)},
		{AvailableFrom: time.Date(2024, time.October, 2, 9, 0, 0, 0, time.Local), AvailableTo: time.Date(2024, time.October, 2, 15, 0, 0, 0, time.Local)},
		{AvailableFrom: time.Date(2024, time.October, 3, 9, 0, 0, 0, time.Local), AvailableTo: time.Date(2024, time.October, 3, 15, 0, 0, 0, time.Local)},
	}
}

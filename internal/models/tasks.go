package models

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	StartTime time.Time
	EndTime   time.Time
	Detail    string
}

type Tasklist []Task

func (t *Task) Date() string {
	yr, month, day := t.StartTime.Date()
	return fmt.Sprintf("%v-%v-%v", day, month, yr)
}

func (t *Task) WeekDayName() string {
	return t.StartTime.Weekday().String()
}

func (t *Task) MonthName() string {
	return t.StartTime.Month().String()
}

func (t *Task) Month() int {
	return int(t.StartTime.Month())
}

func (t *Task) Day() int {
	return t.StartTime.Day()
}

func (t *Task) Yr() int {
	return t.StartTime.Year()
}

func (t *Task) FormattedDate() string {
	return fmt.Sprintf("%v, %v %v, %v", t.WeekDayName(), t.MonthName(), t.Day(), t.Yr())
}

func (t *Task) WeekDay() int {
	return int(t.StartTime.Weekday())
}

func (t *Task) TaskDayKey() string {
	return fmt.Sprintf(
		"%v %v %v %v",
		t.StartTime.Weekday().String(),
		t.StartTime.Month().String(),
		t.StartTime.Day(),
		t.StartTime.Year(),
	)
}

type TaskSchedule map[string]Tasklist

func SortTasksByDay(tasks Tasklist, sortTimes bool) (TaskSchedule, error) {
	sortedTasks := make(map[string]Tasklist, 7)
	if len(tasks) == 0 {
		return nil, fmt.Errorf("empty task list")
	}

	for _, task := range tasks {
		// map key is weekDay MonthName day year
		taskDay := task.TaskDayKey()
		if tsk, exists := sortedTasks[taskDay]; !exists {
			sortedTasks[taskDay] = Tasklist{task}
		} else {
			sortedTasks[taskDay] = append(tsk, task)
		}
	}

	if sortTimes {
		var wg sync.WaitGroup
		keys := make([]string, 0, len(sortedTasks))

		for k := range sortedTasks {
			keys = append(keys, k)
		}

		for _, key := range keys {
			wg.Add(1)
			tks := sortedTasks[key]
			go func() {
				defer wg.Done()
				sortByTimes(&tks)
			}()
		}
		wg.Wait()
	}

	return sortedTasks, nil
}

func sortByTimes(ts *Tasklist) {
	// merge sort
}

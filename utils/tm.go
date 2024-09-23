package utils

import "time"

const MONDAYINDX = 1
const SUNDAYINDX = 0

func WeekStart(dt time.Time, mondayStart bool) time.Time {
	var startIndx int
	if mondayStart {
		startIndx = MONDAYINDX
	} else {
		startIndx = SUNDAYINDX
	}
	// 1 - 4 = -3 -> 4 + -3 = 1
	// 1 - 5 = 4 -> 5 - 4 = 1
	// 1 - 6 = 5 -> 6 - 5 = 1
	// 0 - 4 = -4 -> 4 -4 = 0
	// 0 - 5 = -5 -> 5 -5 = 0
	return dt.AddDate(0, 0, startIndx-int(dt.Weekday()))
}

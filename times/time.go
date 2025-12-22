package times

import "time"

type DateField [3]int

var (
	Year  DateField = [3]int{1, 0, 0}
	Month DateField = [3]int{0, 1, 0}
	Day   DateField = [3]int{0, 0, 1}
)

func YearStart(tm time.Time) time.Time {
	return time.Date(tm.Year(), 1, 1, 0, 0, 0, 0, tm.Location())
}

func YearEnd(tm time.Time) time.Time {
	return time.Date(tm.Year(), 12, 31, 23, 59, 59, 999999999, tm.Location())
}

func MonthStart(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), 1, 0, 0, 0, 0, tm.Location())
}

func MonthEnd(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month()+1, 1, 0, 0, 0, 0, tm.Location()).Add(-time.Nanosecond)
}

func DayStart(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, tm.Location())
}

func DayEnd(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), 23, 59, 59, 999999999, tm.Location())
}

func Split[T string | time.Time](begin time.Time, end time.Time, dateField DateField, fn func(time.Time) T) []T {
	result := make([]T, 0)
	for begin.Before(end) || begin.Equal(end) {
		result = append(result, fn(begin))
		begin = begin.AddDate(dateField[0], dateField[1], dateField[2])
	}
	return result
}

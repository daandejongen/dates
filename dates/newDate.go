package dates

import "errors"

func NewDate(dayOfMonth, month, year int) (Date, []error) {
	date := Date{DayOfMonth: dayOfMonth, Month: month, Year: year}
	return date, []error{
		checkDate(date, hasValidDayOfMonth, "day of month"),
		checkDate(date, hasValidMonth, "month"),
		checkDate(date, hasValidYear, "year"),
	}
}

func checkDate(date Date, requirement func(Date) bool, validatedField string) error {
	if !requirement(date) {
		return errors.New("date " + date.ToString() + " has no valid " + validatedField)
	}
	return nil
}

func hasValidDayOfMonth(date Date) bool {
	return 0 < date.DayOfMonth && date.DayOfMonth <= getNumberOfDaysInMonth(date) 
}

func hasValidMonth(date Date) bool {
	return 0 < date.Month && date.Month <= 12
}

func hasValidYear(date Date) bool {
	return 0 < date.Year
}
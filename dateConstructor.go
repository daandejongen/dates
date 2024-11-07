package dates

func NewDate(dayOfMonth, month, year int) Date {
	date := Date{DayOfMonth: dayOfMonth, Month: month, Year: year}
	if !hasValidDayOfMonth(date) || !hasValidMonth(date) || !hasValidYear(date) {
		return Date{}
	}
	return date
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

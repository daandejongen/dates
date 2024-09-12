package dates

func shiftTime(from Date, amount int, shiftOneUnit func(Date) Date) Date {
	walkingDate := from
	for i := 0; i < amount; i++ {
		walkingDate = shiftOneUnit(walkingDate)
	}
	return walkingDate
}

func addDay(from Date) Date {
	if isLastDayOfTheYear(from) {
		return Date{1, 1, from.Year + 1}
	}
	if isLastDayOfTheMonth(from) {
		return Date{1, from.Month + 1, from.Year}
	}
	return Date{from.DayOfMonth + 1, from.Month, from.Year}
}

func subtractDay(from Date) Date {
	if isFirstDayOfTheYear(from) {
		return Date{31, 12, from.Year - 1}
	}
	if isFirstDayOfTheMonth(from) {
		newDayOfMonth := getNumberOfDaysInMonth(subtractMonth(from))
		return Date{newDayOfMonth, from.Month - 1, from.Year}
	}
	return Date{from.DayOfMonth - 1, from.Month, from.Year}
}

func addMonth(from Date) Date {
	newDayOfMonth := from.DayOfMonth
	newMonth := from.Month + 1
	newYear := from.Year
	if isLastMonthOfTheYear(from) {
		newMonth = 1
		newYear++
	}
	if isLastDayOfTheMonth(from) {
		newDayOfMonth = getNumberOfDaysInMonth(Date{1, newMonth, newYear})
	}
	return Date{newDayOfMonth, newMonth, newYear}
}

func subtractMonth(from Date) Date {
	newDayOfMonth := from.DayOfMonth
	newMonth := from.Month - 1
	newYear := from.Year
	if isFirstMonthOfTheYear(from) {
		newMonth = 12
		newYear--
	}
	if isFirstDayOfTheMonth(from) {
		newDayOfMonth = getNumberOfDaysInMonth(Date{1, newMonth, newYear})
	}
	return Date{newDayOfMonth, newMonth, newYear}
}

func getNumberOfDaysInMonth(date Date) int {
	daysInMonthNormalYear := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	daysInMonthLeapYear := []int{0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if isLeap(date.Year) {
		return daysInMonthLeapYear[date.Month]
	}
	return daysInMonthNormalYear[date.Month]
}

func isFirstDayOfTheMonth(date Date) bool {
	return date.DayOfMonth == 1
}

func isLastDayOfTheMonth(date Date) bool {
	return date.DayOfMonth == getNumberOfDaysInMonth(date)
}

func isFirstDayOfTheYear(date Date) bool {
	return date.DayOfMonth == 1 && date.Month == 1
}

func isLastDayOfTheYear(date Date) bool {
	return date.DayOfMonth == 31 && date.Month == 12
}

func isFirstMonthOfTheYear(date Date) bool {
	return date.Month == 1
}

func isLastMonthOfTheYear(date Date) bool {
	return date.Month == 12
}

func isLeap(year int) bool {
	isMultipleOf4 := year%4 == 0
	isMultipleOf100 := year%100 == 0
	isMultipleOf400 := year%400 == 0
	if !isMultipleOf4 {
		return false
	}
	if isMultipleOf400 {
		return true
	}
	if isMultipleOf100 {
		return false
	}
	return true
}

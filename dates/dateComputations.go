package dates

func (date Date) IsBeforeOrOn(other Date) bool {
	if date == other {
		return true
	}
	return date.IsBefore(other)
}

func (date Date) IsBefore(other Date) bool {
	if date.Year != other.Year {
		return date.Year < other.Year
	}
	if date.Month != other.Month {
		return date.Month < other.Month
	}
	return date.DayOfMonth < other.DayOfMonth
}

func (date Date) IsAfterOrOn(other Date) bool {
	if date == other {
		return true
	}
	return date.IsAfter(other)
}

func (date Date) IsAfter(other Date) bool {
	if date == other {
		return false
	}
	return !date.IsBefore(other)
}

func (date Date) AddDays(amount int) Date {
	return shiftTime(date, amount, addDay)
}

func (date Date) AddMonths(amount int) Date {
	return shiftTime(date, amount, addMonth)
}

func (date Date) SubtractDays(amount int) Date {
	return shiftTime(date, amount, subtractDay)
}

func (date Date) SubtractMonths(amount int) Date {
	return shiftTime(date, amount, subtractMonth)
}
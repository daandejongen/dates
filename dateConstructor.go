package dates

import (
	"github.com/daandejongen/utils/validation"
)

type DateConstructor interface {
	New(dayOfMonth, month, year int) Date
}

type dateConstructor struct {
	validator validation.ObjectValidator[Date]
}

func NewDateConstructor(validator validation.ObjectValidator[Date]) DateConstructor {
	return dateConstructor{validator}
}

func (constructor dateConstructor) New(dayOfMonth, month, year int) Date {
	return constructor.validator.Validate(Date{dayOfMonth, month, year}, []func(Date) bool{
		hasValidDayOfMonth,
		hasValidMonth,
		hasValidYear,
	})
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

package periods

import (
	"github.com/daandejongen/utils/validation"
	"github.com/daandejongen/dates"
)

type PeriodConstructor interface {
	New(start, end dates.Date) Period
}

type periodConstructor struct {
	settings PeriodSettings
	validator validation.ObjectValidator[Period]
}

func NewPeriodConstructor(validator validation.ObjectValidator[Period], settings PeriodSettings) periodConstructor {
	return periodConstructor{settings, validator}
}

func (constructor periodConstructor) New(start, end dates.Date) Period {
	requirements := []func(Period) bool{
		startIsBeforeEnd,
	}
	return constructor.validator.Validate(Period{start, end}, requirements)
}

func (constructor periodConstructor) NewMonth(start dates.Date) Period {
	return constructor.New(start, start.AddMonths(1).SubtractDays(1))
}

func startIsBeforeEnd(period Period) bool {
	return period.Start.IsBefore(period.End)
}

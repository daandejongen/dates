package periods

import (
	"github.com/daandejongen/dates/dates"
)

type PeriodConstructor interface {
	New(start, end dates.Date) Period
}

type periodConstructor struct {
	settings PeriodSettings
}

func NewPeriodConstructor(settings PeriodSettings) periodConstructor {
	return periodConstructor{settings}
}

func (constructor periodConstructor) New(start, end dates.Date) Period {
	period := Period{start, end}
	if !period.Start.IsBefore(period.End) {
		return Period{}
	}
	return period
}

func (constructor periodConstructor) NewMonth(start dates.Date) Period {
	return constructor.New(start, start.AddMonths(1).SubtractDays(1))
}

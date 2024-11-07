package periods

import (
	"github.com/daandejongen/dates/dates"
)

type Period struct {
	Start dates.Date
	End dates.Date
}

func (period Period) Contains(date dates.Date) bool {
	return date.IsAfter(period.Start) && date.IsBefore(period.End)
}
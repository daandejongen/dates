package dates

import (
	"strconv"
	"github.com/daandejongen/utils/formatters"
)

type DateFormatter interface {
	ToDdMmYyyy(Date) string
	ToDdMmYy(Date) string
}

type dateFormatter struct{}

func NewDateFormatter() DateFormatter {
	return dateFormatter{}
}

func (dateFormatter) ToDdMmYyyy(date Date) string {
	nf := formatters.NewNumberFormatter()
	return nf.ToTwoDigits(date.DayOfMonth) + "-" + nf.ToTwoDigits(date.Month) + "-" + strconv.Itoa(date.Year)
}

func (dateFormatter) ToDdMmYy(date Date) string {
	nf := formatters.NewNumberFormatter()
	return nf.ToTwoDigits(date.DayOfMonth) + "-" + nf.ToTwoDigits(date.Month) + "-" + nf.ToTwoDigits(date.Year % 100)
}
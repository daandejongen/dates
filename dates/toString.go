package dates

import (
	"strconv"
	"strings"

	"github.com/daandejongen/utils/formatters"
)

func (date Date) ToString() string {
	nf := formatters.NewNumberFormatter()
	return nf.ToTwoDigits(date.DayOfMonth) + "-" + nf.ToTwoDigits(date.Month) + "-" + strconv.Itoa(date.Year)
}

func (date Date) GetMonthName(language string, abbreviated bool) string {
	month := months.Where(func(month month) bool { return month.number == date.Month })[0]
	if strings.ToLower(language) == "dutch" && abbreviated {
		return month.nameDutchAbbreviated
	}
	if strings.ToLower(language) == "dutch" && !abbreviated {
		return month.nameDutch
	}
	if strings.ToLower(language) == "english" && abbreviated {
		return month.nameEnglishAbbreviated
	}
	return month.nameEnglishAbbreviated
}

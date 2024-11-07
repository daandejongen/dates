package dates

import (
	"github.com/daandejongen/utils/chainable"
)

type month struct {
	number                 int
	daysInMonth            int
	daysInMonthLeapYear    int
	nameDutch              string
	nameDutchAbbreviated   string
	nameEnglish            string
	nameEnglishAbbreviated string
}

var months = chainable.Chainable[month]{
	month{
		number:                 1,
		daysInMonth:            31,
		daysInMonthLeapYear:    31,
		nameDutch:              "januari",
		nameDutchAbbreviated:   "jan",
		nameEnglish:            "January",
		nameEnglishAbbreviated: "Jan",
	},
	month{
		number:                 2,
		daysInMonth:            28,
		daysInMonthLeapYear:    29,
		nameDutch:              "februari",
		nameDutchAbbreviated:   "feb",
		nameEnglish:            "February",
		nameEnglishAbbreviated: "Feb",
	},
	month{
		number:                 3,
		daysInMonth:            31,
		daysInMonthLeapYear:    31,
		nameDutch:              "maart",
		nameDutchAbbreviated:   "mrt",
		nameEnglish:            "March",
		nameEnglishAbbreviated: "Mar",
	},
	month{
		number:                 4,
		daysInMonth:            30,
		daysInMonthLeapYear:    30,
		nameDutch:              "april",
		nameDutchAbbreviated:   "apr",
		nameEnglish:            "April",
		nameEnglishAbbreviated: "Apr",
	},
	month{
		number:                 5,
		daysInMonth:            31,
		daysInMonthLeapYear:    31,
		nameDutch:              "mei",
		nameDutchAbbreviated:   "mei",
		nameEnglish:            "May",
		nameEnglishAbbreviated: "May",
	},
	month{
		number:                 6,
		daysInMonth:            30,
		daysInMonthLeapYear:    30,
		nameDutch:              "juni",
		nameDutchAbbreviated:   "jun",
		nameEnglish:            "June",
		nameEnglishAbbreviated: "Jun",
	},
	month{
		number:                 7,
		daysInMonth:            31,
		daysInMonthLeapYear:    31,
		nameDutch:              "juli",
		nameDutchAbbreviated:   "jul",
		nameEnglish:            "July",
		nameEnglishAbbreviated: "Jul",
	},
	month{
		number:                 8,
		daysInMonth:            31,
		daysInMonthLeapYear:    31,
		nameDutch:              "augustus",
		nameDutchAbbreviated:   "aug",
		nameEnglish:            "August",
		nameEnglishAbbreviated: "Aug",
	},
	month{
		number:                 9,
		daysInMonth:            30,
		daysInMonthLeapYear:    30,
		nameDutch:              "september",
		nameDutchAbbreviated:   "sep",
		nameEnglish:            "September",
		nameEnglishAbbreviated: "Sep",
	},
	month{
		number:                 10,
		daysInMonth:            31,
		daysInMonthLeapYear:    31,
		nameDutch:              "oktober",
		nameDutchAbbreviated:   "okt",
		nameEnglish:            "October",
		nameEnglishAbbreviated: "Oct",
	},
	month{
		number:                 11,
		daysInMonth:            30,
		daysInMonthLeapYear:    30,
		nameDutch:              "november",
		nameDutchAbbreviated:   "nov",
		nameEnglish:            "November",
		nameEnglishAbbreviated: "Jan",
	},
	month{
		number:                 12,
		daysInMonth:            31,
		daysInMonthLeapYear:    31,
		nameDutch:              "december",
		nameDutchAbbreviated:   "dec",
		nameEnglish:            "December",
		nameEnglishAbbreviated: "Dec",
	},
}
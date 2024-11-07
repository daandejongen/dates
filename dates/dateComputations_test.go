package dates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dateIsBeforeTests = []struct {
	this     Date
	other    Date
	expected bool
}{
	{Date{21, 10, 2015}, Date{21, 10, 2015}, false},
	{Date{2, 4, 2023}, Date{1, 3, 2022}, false},
	{Date{1, 1, 2023}, Date{3, 1, 2023}, true},
	{Date{4, 11, 2000}, Date{4, 12, 2000}, true},
}

var AddDaysTests = []struct{
	from Date
	amount int
	expected Date
}{
	{Date{1, 3, 2000}, 6, Date{7, 3, 2000}},
	{Date{31, 7, 2000}, 3, Date{3, 8, 2000}},
	{Date{4, 11, 2000}, 40, Date{14, 12, 2000}},
	{Date{27, 2, 2017}, 4, Date{3, 3, 2017}},
	{Date{29, 12, 1999}, 3, Date{1, 1, 2000}},
}

func TestDateIsBefore(t *testing.T) {
	for _, test := range dateIsBeforeTests {
		t.Run("test", func(t *testing.T) {
			actual := test.this.IsBefore(test.other)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestAddDays(t *testing.T) {
	for _, test := range AddDaysTests {
		assert.Equal(t, test.expected, test.from.AddDays(test.amount))
	}
}
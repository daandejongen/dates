package dates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dateToStringTests = []struct {
	date Date
	expected string
}{
	{Date{1, 1, 2000}, "01-01-2000"},
	{Date{19, 11, 2020}, "19-11-2020"},
}

func TestDateToString(t *testing.T) {
	for _, test := range dateToStringTests {
		assert.Equal(t, test.expected, test.date.ToString())
	}
}

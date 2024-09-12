package dates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToDdMmYyyy(t *testing.T) {
	df := NewDateFormatter()
	assert.Equal(t, "01-01-2000", df.ToDdMmYyyy(Date{1, 1, 2000}))
	assert.Equal(t, "11-12-2024", df.ToDdMmYyyy(Date{11, 12, 2024}))
}

func TestToDdMmYy(t *testing.T) {
	df := NewDateFormatter()
	assert.Equal(t, "01-01-00", df.ToDdMmYy(Date{1, 1, 2000}))
	assert.Equal(t, "11-12-24", df.ToDdMmYy(Date{11, 12, 2024}))
	assert.Equal(t, "11-12-96", df.ToDdMmYy(Date{11, 12, 1896}))
}

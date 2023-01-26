package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScheduleConverter_NoError(t *testing.T) {
	result, err := ScheduleConverter("2023-01-26")

	assert.NoError(t, err)
	assert.Equal(t, time.Time(time.Date(2023, time.January, 26, 0, 0, 0, 0, time.UTC)), result)
}

func TestScheduleConverter_Error(t *testing.T) {
	result, err := ScheduleConverter("2023-02-30")

	assert.NotNil(t, err)
	assert.Equal(t, time.Time(time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)), result)
}

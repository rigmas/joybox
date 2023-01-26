package utils

import (
	"fmt"
	"time"
)

const (
	timeLayout = "2006-01-02"
)

func ScheduleConverter(pickupSchedule string) (time.Time, error) {
	timeSchedule, err := time.Parse(timeLayout, pickupSchedule)
	if err != nil {
		return time.Time{}, fmt.Errorf("Invalid date: %v", pickupSchedule)
	}

	return timeSchedule, nil
}

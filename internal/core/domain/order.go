package domain

import "time"

type Order struct {
	Book           Book
	PickupSchedule time.Time
}

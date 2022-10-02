package models

import (
	"time"
)

type Reservation struct {
	ID        string
	ServiceID string
	Date      time.Time
	Comment   *string
	Approved  bool
}

type ReservationUpdate struct {
	Date     *time.Time
	Comment  **string
	Approved *bool
}

func (r *ReservationUpdate) IsEmpty() bool {
	return r.Date == nil && r.Comment == nil && r.Approved == nil
}

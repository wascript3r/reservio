package models

import (
	"time"

	cmodels "github.com/wascript3r/reservio/pkg/features/client/models"
	smodels "github.com/wascript3r/reservio/pkg/features/service/models"
)

type Reservation struct {
	ID        string
	ServiceID string
	ClientID  string
	Date      time.Time
	Comment   *string
	Approved  bool
}

type FullReservation struct {
	ID        string
	ServiceID string
	Client    *cmodels.ClientInfo
	Date      time.Time
	Comment   *string
	Approved  bool
}

type ClientReservation struct {
	ID       string
	Service  *smodels.FullService
	Date     time.Time
	Comment  *string
	Approved bool
}

type ReservationUpdate struct {
	Date     *time.Time
	Comment  **string
	Approved *bool
}

func (r *ReservationUpdate) IsEmpty() bool {
	return r.Date == nil && r.Comment == nil && r.Approved == nil
}

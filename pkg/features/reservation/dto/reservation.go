package dto

import (
	"html"

	cdto "github.com/wascript3r/reservio/pkg/features/client/dto"
	sdto "github.com/wascript3r/reservio/pkg/features/service/dto"
)

// Create

type CreateReq struct {
	sdto.ServiceReq
	ClientID string  `json:"-" validate:"required,uuid"`
	Date     string  `json:"date" validate:"required,r_date"`
	Comment  *string `json:"comment" validate:"omitempty,r_comment"`
}

func (c *CreateReq) Escape() {
	if c.Comment != nil {
		*c.Comment = html.EscapeString(*c.Comment)
	}
}

type CreateRes struct {
	ID        string  `json:"id"`
	ServiceID string  `json:"serviceID"`
	ClientID  string  `json:"clientID"`
	Date      string  `json:"date"`
	Comment   *string `json:"comment"`
}

// Get

type ReservationReq struct {
	sdto.ServiceReq
	ReservationID string `json:"-" validate:"required,uuid"`
}

type GetReq struct {
	ReservationReq
	ClientID *string `json:"-" validate:"omitempty,uuid"`
}

type Reservation struct {
	ID        string       `json:"id"`
	ServiceID string       `json:"serviceID"`
	Client    *cdto.Client `json:"client"`
	Date      string       `json:"date"`
	Comment   *string      `json:"comment"`
}

type ReservationMeta struct {
	ID        string `json:"id"`
	ServiceID string `json:"serviceID"`
	Date      string `json:"date"`
}

type GetRes Reservation

// GetAll

type GetAllReq sdto.ServiceReq

type GetAllRes struct {
	Reservations []*Reservation `json:"reservations"`
}

// GetAllMeta

type GetAllMetaReq sdto.ServiceReq

type GetAllMetaRes struct {
	Reservations []*ReservationMeta `json:"reservations"`
}

// GetAllByClient

type ClientReservation struct {
	ID      string            `json:"id"`
	Service *sdto.FullService `json:"service"`
	Date    string            `json:"date"`
	Comment *string           `json:"comment"`
}

type GetAllByClientReq cdto.ClientReq

type GetAllByClientRes struct {
	Reservations []*ClientReservation `json:"reservations"`
}

// Update

type Comment struct {
	Value *string `json:"value" validate:"omitempty,r_comment"`
}

type UpdateReq struct {
	ReservationReq
	ClientID string   `json:"-" validate:"required,uuid"`
	Date     *string  `json:"date" validate:"omitempty,r_date"`
	Comment  *Comment `json:"comment" validate:"omitempty,dive"`
}

func (u *UpdateReq) Escape() {
	if u.Comment != nil && u.Comment.Value != nil {
		*u.Comment.Value = html.EscapeString(*u.Comment.Value)
	}
}

// Delete

type DeleteReq struct {
	ReservationReq
	ClientID string `json:"-" validate:"required,uuid"`
}

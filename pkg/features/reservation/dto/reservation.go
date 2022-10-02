package dto

import (
	"html"

	sdto "github.com/wascript3r/reservio/pkg/features/service/dto"
)

// Create

type CreateReq struct {
	sdto.ServiceReq
	Date    string  `json:"date" validate:"required,r_date"`
	Comment *string `json:"comment" validate:"omitempty,r_comment"`
}

func (c *CreateReq) Escape() {
	if c.Comment != nil {
		*c.Comment = html.EscapeString(*c.Comment)
	}
}

type CreateRes struct {
	ID string `json:"id"`
}

// Get

type ReservationReq struct {
	sdto.ServiceReq
	ReservationID string `json:"-" validate:"required,uuid"`
}

type GetReq ReservationReq

type Reservation struct {
	ID        string  `json:"id"`
	ServiceID string  `json:"serviceID"`
	Date      string  `json:"date"`
	Comment   *string `json:"comment"`
}

type GetRes Reservation

// GetAll

type GetAllReq sdto.ServiceReq

type GetAllRes struct {
	Reservations []*Reservation `json:"reservations"`
}

// Update

type UpdateReq struct {
	ReservationReq
	Date    *string `json:"date" validate:"omitempty,r_date"`
	Comment *struct {
		Value *string `json:"value" validate:"omitempty,r_comment"`
	} `json:"comment" validate:"omitempty,dive"`
}

func (u *UpdateReq) Escape() {
	if u.Comment != nil && u.Comment.Value != nil {
		*u.Comment.Value = html.EscapeString(*u.Comment.Value)
	}
}

// Delete

type DeleteReq ReservationReq

package dto

import (
	"html"

	cdto "github.com/wascript3r/reservio/pkg/features/company/dto"
	"github.com/wascript3r/reservio/pkg/features/service/models"
)

// Create

type WorkHours struct {
	From string `json:"from" validate:"required,s_time"`
	To   string `json:"to" validate:"required,s_time"`
}

type WorkSchedule map[models.Weekday]*WorkHours

type CreateReq struct {
	cdto.CompanyReq
	Title           string       `json:"title" validate:"required,s_title"`
	Description     string       `json:"description" validate:"required,s_description"`
	SpecialistName  *string      `json:"specialistName" validate:"omitempty,s_specialist_name"`
	SpecialistPhone *string      `json:"specialistPhone" validate:"omitempty,s_phone"`
	VisitDuration   int          `json:"visitDuration" validate:"required,gt=0"`
	WorkSchedule    WorkSchedule `json:"workSchedule" validate:"required,gt=0,s_work_schedule,dive,required"`
}

func (c *CreateReq) Escape() {
	c.Title = html.EscapeString(c.Title)
	c.Description = html.EscapeString(c.Description)
	if c.SpecialistName != nil {
		*c.SpecialistName = html.EscapeString(*c.SpecialistName)
	}
	if c.SpecialistPhone != nil {
		*c.SpecialistPhone = html.EscapeString(*c.SpecialistPhone)
	}
}

type Service struct {
	ID              string       `json:"id"`
	CompanyID       string       `json:"companyID"`
	Title           string       `json:"title"`
	Description     string       `json:"description"`
	SpecialistName  *string      `json:"specialistName"`
	SpecialistPhone *string      `json:"specialistPhone"`
	VisitDuration   int          `json:"visitDuration"`
	WorkSchedule    WorkSchedule `json:"workSchedule"`
}

type CreateRes Service

// Get

type ServiceReq struct {
	cdto.CompanyReq
	ServiceID string `json:"-" validate:"required,uuid"`
}

type GetReq ServiceReq

type FullService struct {
	ID              string        `json:"id"`
	Company         *cdto.Company `json:"company"`
	Title           string        `json:"title"`
	Description     string        `json:"description"`
	SpecialistName  *string       `json:"specialistName"`
	SpecialistPhone *string       `json:"specialistPhone"`
	VisitDuration   int           `json:"visitDuration"`
	WorkSchedule    WorkSchedule  `json:"workSchedule"`
}

type GetRes Service

// GetAll

type GetAllReq cdto.CompanyReq

type GetAllRes struct {
	Services []*Service `json:"services"`
}

// Update

type UpdateReq struct {
	ServiceReq
	Title          *string `json:"title" validate:"omitempty,s_title"`
	Description    *string `json:"description" validate:"omitempty,s_description"`
	SpecialistName *struct {
		Value *string `json:"value" validate:"omitempty,s_specialist_name"`
	} `json:"specialistName" validate:"omitempty,dive"`
	SpecialistPhone *struct {
		Value *string `json:"value" validate:"omitempty,s_phone"`
	} `json:"specialistPhone" validate:"omitempty,dive"`
	WorkSchedule *WorkSchedule `json:"workSchedule" validate:"omitempty,gt=0,s_work_schedule,dive,required"`
}

func (u *UpdateReq) Escape() {
	if u.Title != nil {
		*u.Title = html.EscapeString(*u.Title)
	}
	if u.Description != nil {
		*u.Description = html.EscapeString(*u.Description)
	}
	if u.SpecialistName != nil && u.SpecialistName.Value != nil {
		*u.SpecialistName.Value = html.EscapeString(*u.SpecialistName.Value)
	}
	if u.SpecialistPhone != nil && u.SpecialistPhone.Value != nil {
		*u.SpecialistPhone.Value = html.EscapeString(*u.SpecialistPhone.Value)
	}
}

// Delete

type DeleteReq ServiceReq

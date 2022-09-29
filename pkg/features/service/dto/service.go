package dto

import (
	cdto "github.com/wascript3r/reservio/pkg/features/company/dto"
	"github.com/wascript3r/reservio/pkg/features/service/models"
)

// Create

type WorkHours struct {
	From string `json:"from" validate:"required,s_time"`
	To   string `json:"to" validate:"required,s_time"`
}

type WorkSchedule map[models.Weekday]WorkHours

type CreateReq struct {
	cdto.CompanyReq
	Title           string       `json:"title" validate:"required,s_title"`
	Description     string       `json:"description" validate:"required,s_description"`
	SpecialistName  *string      `json:"specialistName" validate:"omitempty,s_specialist_name"`
	SpecialistPhone *string      `json:"specialistPhone" validate:"omitempty,s_phone"`
	WorkSchedule    WorkSchedule `json:"workSchedule" validate:"required,gt=0,s_work_schedule,dive,required"`
}

type CreateRes struct {
	ID string `json:"id"`
}

package models

import (
	"errors"
)

var ErrInvalidWeekday = errors.New("invalid weekday")

type Weekday string

const (
	Monday    Weekday = "monday"
	Tuesday   Weekday = "tuesday"
	Wednesday Weekday = "wednesday"
	Thursday  Weekday = "thursday"
	Friday    Weekday = "friday"
	Saturday  Weekday = "saturday"
	Sunday    Weekday = "sunday"
)

var weekdays = map[int]Weekday{
	0: Sunday,
	1: Monday,
	2: Tuesday,
	3: Wednesday,
	4: Thursday,
	5: Friday,
	6: Saturday,
}

func NewWeekday(d int) (Weekday, error) {
	if w, ok := weekdays[d]; ok {
		return w, nil
	}
	return "", ErrInvalidWeekday
}

func (w Weekday) IsValid() bool {
	switch w {
	case Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday:
		return true
	default:
		return false
	}
}

type WorkHours struct {
	From string
	To   string
}

type WorkSchedule map[Weekday]*WorkHours

type Service struct {
	ID              string
	CompanyID       string
	Title           string
	Description     string
	SpecialistName  *string
	SpecialistPhone *string
	VisitDuration   int
	WorkSchedule    WorkSchedule
}

type ServiceUpdate struct {
	Title           *string
	Description     *string
	SpecialistName  **string
	SpecialistPhone **string
	WorkSchedule    *WorkSchedule
}

func (s *ServiceUpdate) IsEmpty() bool {
	return s.Title == nil && s.Description == nil && s.SpecialistName == nil &&
		s.SpecialistPhone == nil && s.WorkSchedule == nil
}

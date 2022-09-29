package models

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
	WorkSchedule    WorkSchedule
}

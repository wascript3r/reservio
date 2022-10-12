package validator

import (
	"time"

	"github.com/wascript3r/reservio/pkg/features/reservation"

	gvalidator "github.com/go-playground/validator/v10"
	"github.com/wascript3r/reservio/pkg/features/service/models"
	"github.com/wascript3r/reservio/pkg/validator"
	"github.com/wascript3r/reservio/pkg/validator/gov"
)

const timeFormat = "15:04"

type Validator struct {
	govalidate *gvalidator.Validate
}

func New() *Validator {
	v := gvalidator.New()

	r := newRules()
	r.attachTo(v)

	return &Validator{v}
}

func (v *Validator) RawRequest(s any) validator.Error {
	err := v.govalidate.Struct(s)
	if err != nil {
		return gov.Translate(err)
	}
	return nil
}

func (v *Validator) ReservationDate(ss *models.Service, date time.Time) error {
	if date.Before(time.Now()) {
		return reservation.DateIsInPastError
	}

	weekday, err := models.NewWeekday(int(date.Weekday()))
	if err != nil {
		return err
	}
	workHours, ok := ss.WorkSchedule[weekday]
	if !ok {
		return reservation.ServiceNotAvailableError
	}
	from, err := time.Parse(timeFormat, workHours.From)
	if err != nil {
		return err
	}
	to, err := time.Parse(timeFormat, workHours.To)
	if err != nil {
		return err
	}
	fromDate := time.Date(date.Year(), date.Month(), date.Day(), from.Hour(), from.Minute(), 0, 0, time.UTC)
	toDate := time.Date(date.Year(), date.Month(), date.Day(), to.Hour(), to.Minute(), 0, 0, time.UTC)

	endDate := date.Add(time.Duration(ss.VisitDuration) * time.Minute)
	if date.Before(fromDate) || endDate.After(toDate) {
		return reservation.ServiceNotAvailableError
	}

	for !fromDate.After(date) {
		if fromDate.Equal(date) {
			return nil
		}
		fromDate = fromDate.Add(time.Duration(ss.VisitDuration) * time.Minute)
	}

	return reservation.InvalidTimeError
}

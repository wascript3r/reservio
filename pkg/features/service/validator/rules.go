package validator

import (
	"reflect"
	"time"

	"github.com/wascript3r/reservio/pkg/features/service/models"

	gvalidator "github.com/go-playground/validator/v10"
	"github.com/wascript3r/reservio/pkg/validator/gov"
)

const timeFormat = "15:04"

type rules struct {
	aliases []gov.AliasRule
	fns     []gov.FnRule
}

func newRules() rules {
	return rules{
		aliases: []gov.AliasRule{
			gov.NewAliasRule("s_title", "gte=3,lte=100"),
			gov.NewAliasRule("s_description", "gte=5"),
			gov.NewAliasRule("s_specialist_name", "gte=5,lte=100"),
			gov.NewAliasRule("s_phone", "e164"),
		},
		fns: []gov.FnRule{
			gov.NewFnRule("s_work_schedule", validateWorkSchedule),
			gov.NewFnRule("s_time", validateTime),
		},
	}
}

func (r rules) attachTo(v *gvalidator.Validate) {
	for _, a := range r.aliases {
		v.RegisterAlias(a.Alias(), a.Tags())
	}
	for _, f := range r.fns {
		v.RegisterValidation(f.Name(), f.IsValid)
	}
}

func validateTime(fl gvalidator.FieldLevel) bool {
	field := fl.Field()

	if field.Kind() != reflect.String {
		return false
	}
	t := field.String()

	_, err := time.Parse(timeFormat, t)
	return err == nil
}

func validateWorkSchedule(fl gvalidator.FieldLevel) bool {
	field := fl.Field()

	if field.Kind() != reflect.Map {
		return false
	}

	var weekDay models.Weekday
	for _, key := range field.MapKeys() {
		if key.Kind() != reflect.String {
			return false
		}
		weekDay = models.Weekday(key.String())
		if !weekDay.IsValid() {
			return false
		}
	}

	return true
}

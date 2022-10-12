package reservation

import (
	"time"

	"github.com/wascript3r/reservio/pkg/features/service/models"
	"github.com/wascript3r/reservio/pkg/validator"
)

type Validator interface {
	RawRequest(s any) validator.Error
	ReservationDate(ss *models.Service, date time.Time) error
}

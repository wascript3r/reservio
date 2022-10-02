package reservation

import (
	"context"
	"time"

	"github.com/wascript3r/reservio/pkg/features/reservation/models"
)

type Repository interface {
	Insert(ctx context.Context, rs *models.Reservation) (id string, err error)
	Get(ctx context.Context, companyID, serviceID, reservationID string, onlyApprovedCompany bool) (*models.Reservation, error)
	GetAll(ctx context.Context, companyID, serviceID string, onlyApprovedCompany bool) ([]*models.Reservation, error)
	Update(ctx context.Context, companyID, serviceID, reservationID string, ru *models.ReservationUpdate) error
	Delete(ctx context.Context, companyID, serviceID, reservationID string) error
	Exists(ctx context.Context, companyID, serviceID string, date time.Time) (bool, error)
}

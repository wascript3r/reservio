package models

import (
	"time"

	umodels "github.com/wascript3r/reservio/pkg/features/user/models"
)

type RefreshToken struct {
	ID        string
	UserID    string
	ExpiresAt time.Time
}

type Claims struct {
	UserID string
	Role   umodels.Role
}

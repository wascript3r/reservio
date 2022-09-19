package usecase

import (
	"context"
	"html"
	"time"

	"github.com/wascript3r/reservio/pkg/features/user"
	"github.com/wascript3r/reservio/pkg/features/user/dto"
	"github.com/wascript3r/reservio/pkg/features/user/models"
)

type Usecase struct {
	userRepo   user.Repository
	ctxTimeout time.Duration

	pwHasher  user.PwHasher
	validator user.Validator
}

func New(ur user.Repository, t time.Duration, ph user.PwHasher, v user.Validator) *Usecase {
	return &Usecase{
		userRepo:   ur,
		ctxTimeout: t,

		pwHasher:  ph,
		validator: v,
	}
}

func (u *Usecase) Create(ctx context.Context, req *dto.CreateReq) error {
	if err := u.validator.RawRequest(req); err != nil {
		return user.InvalidInputError
	}

	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	var err error

	err = u.validator.EmailUniqueness(c, req.Email)
	if err != nil {
		if err == user.ErrEmailExists {
			return user.EmailAlreadyExistsError
		}
		return err
	}

	hash, err := u.pwHasher.Hash(req.Password)
	if err != nil {
		return err
	}

	us := &models.User{
		Email:    html.EscapeString(req.Email),
		Password: hash,
		Role:     models.ClientRole,
	}

	_, err = u.userRepo.Insert(c, us)
	if err != nil {
		return err
	}

	return nil
}

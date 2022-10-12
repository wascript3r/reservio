package usecase

import (
	"context"
	"time"

	tdto "github.com/wascript3r/reservio/pkg/features/token/dto"

	"github.com/wascript3r/reservio/pkg/features/token"
	"github.com/wascript3r/reservio/pkg/features/user"
	"github.com/wascript3r/reservio/pkg/features/user/dto"
	"github.com/wascript3r/reservio/pkg/features/user/models"
	"github.com/wascript3r/reservio/pkg/repository"
)

type Usecase struct {
	userRepo   user.Repository
	ctxTimeout time.Duration

	pwHasher   user.PwHasher
	tokenUcase token.Usecase
	validator  user.Validator
}

func New(ur user.Repository, t time.Duration, ph user.PwHasher, tu token.Usecase, v user.Validator) *Usecase {
	return &Usecase{
		userRepo:   ur,
		ctxTimeout: t,

		pwHasher:   ph,
		tokenUcase: tu,
		validator:  v,
	}
}

func (u *Usecase) Create(ctx context.Context, req *dto.CreateReq, role models.Role, validateReq bool) (string, error) {
	if validateReq {
		if err := u.validator.RawRequest(req); err != nil {
			return "", user.InvalidInputError
		}
	}

	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	var err error

	err = u.validator.EmailUniqueness(c, req.Email)
	if err != nil {
		return "", err
	}

	hash, err := u.pwHasher.Hash(req.Password)
	if err != nil {
		return "", err
	}

	req.Escape()
	us := &models.User{
		Email:    req.Email,
		Password: hash,
		Role:     role,
	}

	return u.userRepo.Insert(c, us)
}

func (u *Usecase) Authenticate(ctx context.Context, req *dto.AuthenticateReq) (*tdto.TokenPair, error) {
	if err := u.validator.RawRequest(req); err != nil {
		return nil, user.InvalidInputError
	}

	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	us, err := u.userRepo.GetByEmail(c, req.Email)
	if err != nil {
		if err == repository.ErrNoItems {
			return nil, user.InvalidCredentialsError
		}
		return nil, err
	}

	if err := u.pwHasher.Validate(us.Password, req.Password); err != nil {
		return nil, user.InvalidCredentialsError
	}

	return u.tokenUcase.IssuePair(ctx, us)
}

func (u *Usecase) Logout(ctx context.Context) error {
	claims, err := u.tokenUcase.LoadCtx(ctx)
	if err != nil {
		return err
	}

	return u.tokenUcase.RevokeRefresh(ctx, claims.RefreshTokenID)
}

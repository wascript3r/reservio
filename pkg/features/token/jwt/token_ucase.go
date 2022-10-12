package jwt

import (
	"context"
	"errors"
	"time"

	"github.com/wascript3r/reservio/pkg/repository"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/wascript3r/reservio/pkg/features/token"
	"github.com/wascript3r/reservio/pkg/features/token/dto"
	"github.com/wascript3r/reservio/pkg/features/token/models"
	umodels "github.com/wascript3r/reservio/pkg/features/user/models"
)

type ctxKey struct{}

var (
	signingMethod = jwt.SigningMethodHS256

	ErrInvalidTokenAlg = errors.New("invalid token algorithm")
)

type Options struct {
	PrivateKey        []byte
	AccessExpiration  time.Duration
	RefreshExpiration time.Duration
	Issuer            string
}

type accessClaims struct {
	jwt.StandardClaims
	*dto.AccessClaims
}

type refreshClaims struct {
	jwt.StandardClaims
	*dto.RefreshClaims
}

type Usecase struct {
	*Options
	tx         repository.Transactor
	tokenRepo  token.Repository
	ctxTimeout time.Duration

	validator token.Validator
}

func New(o *Options, tx repository.Transactor, tr token.Repository, t time.Duration, v token.Validator) *Usecase {
	return &Usecase{
		Options:    o,
		tx:         tx,
		tokenRepo:  tr,
		ctxTimeout: t,

		validator: v,
	}
}

func (u *Usecase) buildStandardClaims(exp time.Duration) jwt.StandardClaims {
	return jwt.StandardClaims{
		ExpiresAt: jwt.At(time.Now().Add(exp)),
		// ID:        "",
		IssuedAt: jwt.Now(),
		Issuer:   u.Issuer,
		// NotBefore: nil,
		// Subject:   "",
	}
}

func (u *Usecase) generate(claims jwt.Claims) (string, error) {
	t := jwt.NewWithClaims(signingMethod, claims)
	return t.SignedString(u.PrivateKey)
}

func (u *Usecase) generateAccess(claims *dto.AccessClaims) (string, error) {
	combined := &accessClaims{
		StandardClaims: u.buildStandardClaims(u.AccessExpiration),
		AccessClaims:   claims,
	}
	return u.generate(combined)
}

func (u *Usecase) generateRefresh(id string) (string, error) {
	combined := &refreshClaims{
		StandardClaims: u.buildStandardClaims(u.RefreshExpiration),
		RefreshClaims: &dto.RefreshClaims{
			RefreshTokenID: id,
		},
	}
	return u.generate(combined)
}

func (u *Usecase) IssuePair(ctx context.Context, us *umodels.User) (*dto.TokenPair, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	pair := &dto.TokenPair{}
	err := u.tx.WithinTx(c, func(c context.Context) error {
		rID, err := u.tokenRepo.Insert(c, &models.RefreshToken{
			UserID:    us.ID,
			ExpiresAt: time.Now().Add(u.RefreshExpiration),
		})
		if err != nil {
			return err
		}

		pair.RefreshToken, err = u.generateRefresh(rID)
		if err != nil {
			return err
		}

		pair.AccessToken, err = u.generateAccess(&dto.AccessClaims{
			UserID:         us.ID,
			Role:           us.Role,
			RefreshTokenID: rID,
		})
		return err
	})
	if err != nil {
		return nil, err
	}

	return pair, nil
}

func (u *Usecase) RenewAccess(ctx context.Context, req *dto.RenewAccessReq) (*dto.RenewAccessRes, error) {
	if err := u.validator.RawRequest(req); err != nil {
		return nil, token.InvalidInputError.SetData(err.GetData())
	}

	refresh, err := u.ParseRefresh(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	cls, err := u.tokenRepo.GetClaims(c, refresh.RefreshTokenID)
	if err != nil {
		if err == repository.ErrNoItems {
			return nil, token.InvalidOrExpiredTokenError
		}
		return nil, err
	}

	access, err := u.generateAccess(&dto.AccessClaims{
		UserID:         cls.UserID,
		Role:           cls.Role,
		RefreshTokenID: refresh.RefreshTokenID,
	})
	if err != nil {
		return nil, err
	}

	return &dto.RenewAccessRes{
		AccessToken: access,
	}, nil
}

func (u *Usecase) parse(tkn string, claims jwt.Claims) (jwt.Claims, error) {
	t, err := jwt.ParseWithClaims(tkn, claims, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok || t.Method.Alg() != signingMethod.Alg() {
			return nil, ErrInvalidTokenAlg
		}
		return u.PrivateKey, nil
	})
	if err != nil || !t.Valid {
		return nil, token.InvalidOrExpiredTokenError
	}

	return t.Claims, nil
}

func (u *Usecase) ParseAccess(tkn string) (*dto.AccessClaims, error) {
	claims, err := u.parse(tkn, &accessClaims{})
	if err != nil {
		return nil, err
	}

	access, ok := claims.(*accessClaims)
	if !ok || access.AccessClaims == nil {
		return nil, token.FaultyTokenError
	}

	return access.AccessClaims, nil
}

func (u *Usecase) ParseRefresh(tkn string) (*dto.RefreshClaims, error) {
	claims, err := u.parse(tkn, &refreshClaims{})
	if err != nil {
		return nil, err
	}

	refresh, ok := claims.(*refreshClaims)
	if !ok || refresh.RefreshClaims == nil {
		return nil, token.FaultyTokenError
	}

	return refresh.RefreshClaims, nil
}

func (u *Usecase) StoreCtx(ctx context.Context, claims *dto.AccessClaims) context.Context {
	return context.WithValue(ctx, ctxKey{}, claims)
}

func (u *Usecase) LoadCtx(ctx context.Context) (*dto.AccessClaims, error) {
	claims, ok := ctx.Value(ctxKey{}).(*dto.AccessClaims)
	if !ok {
		return nil, token.ErrCannotLoadToken
	}
	return claims, nil
}

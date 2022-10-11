package jwt

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/wascript3r/reservio/pkg/features/token"
	"github.com/wascript3r/reservio/pkg/features/user/models"
)

type ctxKey struct{}

var (
	signingMethod = jwt.SigningMethodHS256

	ErrInvalidTokenAlg    = errors.New("invalid token algorithm")
	ErrInvalidToken       = errors.New("invalid token")
	ErrInvalidTokenClaims = errors.New("invalid token claims")
)

type AuthClaims struct {
	jwt.StandardClaims
	token.UserClaims
}

type Usecase struct {
	privateKey []byte
	expiration time.Duration
	issuer     string
}

func New(privateKey []byte, expiration time.Duration, issuer string) *Usecase {
	return &Usecase{
		privateKey: privateKey,
		expiration: expiration,
		issuer:     issuer,
	}
}

func (u *Usecase) Generate(us *models.User) (string, error) {
	t := jwt.NewWithClaims(signingMethod, AuthClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(u.expiration)),
			// ID:        "",
			IssuedAt: jwt.Now(),
			Issuer:   u.issuer,
			// NotBefore: nil,
			// Subject:   "",
		},
		UserClaims: token.UserClaims{
			UserID: us.ID,
			Role:   us.Role,
		},
	})
	return t.SignedString(u.privateKey)
}

func (u *Usecase) Parse(tkn string) (*token.UserClaims, error) {
	t, err := jwt.ParseWithClaims(tkn, &AuthClaims{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok || t.Method.Alg() != signingMethod.Alg() {
			return nil, ErrInvalidTokenAlg
		}
		return u.privateKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, token.InvalidOrExpiredTokenError
	}

	claims, ok := t.Claims.(*AuthClaims)
	if !ok {
		return nil, ErrInvalidTokenClaims
	}

	return &claims.UserClaims, nil
}

func (u *Usecase) StoreCtx(ctx context.Context, claims *token.UserClaims) context.Context {
	return context.WithValue(ctx, ctxKey{}, claims)
}

func (u *Usecase) LoadCtx(ctx context.Context) (*token.UserClaims, error) {
	claims, ok := ctx.Value(ctxKey{}).(*token.UserClaims)
	if !ok {
		return nil, token.ErrCannotLoadToken
	}
	return claims, nil
}

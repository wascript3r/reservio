package jwt

import (
	"errors"
	"time"

	"github.com/wascript3r/reservio/pkg/features/token"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/wascript3r/reservio/pkg/features/user/models"
)

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
			Role:   us.Role.String(),
		},
	})
	return t.SignedString(u.privateKey)
}

func (u *Usecase) Parse(token string) (*token.UserClaims, error) {
	t, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(t *jwt.Token) (interface{}, error) {
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
		return nil, ErrInvalidToken
	}

	claims, ok := t.Claims.(*AuthClaims)
	if !ok {
		return nil, ErrInvalidTokenClaims
	}

	return &claims.UserClaims, nil
}

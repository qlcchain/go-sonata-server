package auth

import (
	"errors"
	"time"

	"github.com/qlcchain/go-sonata-server/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var (
	ErrInvalidRoleclaims = errors.New("invalid role claims")
)

type RoleClaims struct {
	Roles []string `json:"roles"`
	jwt.StandardClaims
}

func NewRoleClaims(roles []string) *RoleClaims {
	return &RoleClaims{
		Roles: roles,
		StandardClaims: jwt.StandardClaims{
			Audience:  "QLCChain Bot",
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Id:        uuid.New().String(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "QLCChain Bot",
			NotBefore: 0,
			Subject:   "sonata",
		},
	}
}

func (r *RoleClaims) String() string {
	return util.ToString(r)
}

func ParseAndCheckToken(token string, pubKey interface{}) (*RoleClaims, error) {
	if parsedToken, err := jwt.ParseWithClaims(token, &RoleClaims{}, func(parsedToken *jwt.Token) (interface{}, error) {
		return pubKey, nil
	}); err == nil {
		if claims, ok := parsedToken.Claims.(*RoleClaims); ok {
			return claims, nil
		} else {
			return nil, ErrInvalidRoleclaims
		}
	} else {
		return nil, err
	}
}

func NewToken(claim *RoleClaims, private interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES512, claim)
	return token.SignedString(private)
}

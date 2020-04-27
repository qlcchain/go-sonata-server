package auth

import (
	"crypto/rsa"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var (
	// Keys used to sign and verify our tokens
	verifyKey *rsa.PublicKey
)

type RoleClaims struct {
	Roles []string `json:"roles"`
	jwt.StandardClaims
}

func ParseAndCheckToken(token string) (*RoleClaims, error) {
	if parsedToken, err := jwt.ParseWithClaims(token, &RoleClaims{}, func(parsedToken *jwt.Token) (interface{}, error) {
		// the key used to validate tokens
		return verifyKey, nil
	}); err == nil {
		if claims, ok := parsedToken.Claims.(*RoleClaims); ok && parsedToken.Valid {
			return claims, nil
		} else {
			return nil, fmt.Errorf("invalid role claims")
		}
	} else {
		return nil, err
	}
}

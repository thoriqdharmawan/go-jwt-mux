package config

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("qfuiasgifouagsfoiugaoisfug")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}

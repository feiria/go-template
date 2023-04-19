package request

import (
	"github.com/golang-jwt/jwt/v4"
)

type BaseClaims struct {
	Uuid string
	ID   uint
	Name string
}

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

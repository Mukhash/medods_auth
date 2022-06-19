package models

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
	UUID string `bson:"uuid"`
}

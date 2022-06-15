package models

import "github.com/dgrijalva/jwt-go"

var HmacSecretKey string = "1873ff373ae9f39c4f5af53450ebe7d174d3cc3066b5a6557a38728b094e3775"

type Claims struct {
	jwt.StandardClaims
	UUID string `bson:"_id"`
}

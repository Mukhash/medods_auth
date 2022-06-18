package models

import "github.com/golang-jwt/jwt"

var HmacAccessSecret string = "2f6877ee76893342bb0989352baac52a9b3965083958cc006345f4390ded483805bf1cc7bab8e343a9e27f837d8fd43aad75a48d0f0e60f497cadb4176cbe9ed"
var HmacRefreshSecret string = "067bfb574b2995f5205cad66a0d135e34f056b5b28adba656b40922d92fc78ef95021f22a6d179adac5b67ec146b3e48ab882cd6cf336a72c4a5678e3bc9e350"

type Claims struct {
	jwt.StandardClaims
	UUID string `bson:"_id"`
}

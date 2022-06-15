package models

type Token struct {
	Access  string
	Refresh string
}

type User struct {
	UUID         string `bson:"_id"`
	RefreshToken string `bson:"refresh_token"`
}

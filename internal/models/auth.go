package models

type Token struct {
	Access  string `json:"accessToken"`
	Refresh string `json:"refreshToken"`
}

type User struct {
	UUID         string `bson:"uuid"`
	RefreshToken string `bson:"refresh_token"`
}

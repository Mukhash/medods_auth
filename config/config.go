package config

import "time"

type Config struct {
	API API
	DB  DB
}

type API struct {
	Address         string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
	MainPath        string
	TokenTTL        int64
	JwtKey          string
}

type DB struct {
	URL          string
	SchemaName   string
	MaxOpenConns int
	MaxIdleConns int
}

func DefaultConfig() *Config {
	return &Config{
		API: API{
			Address:         ":5000",
			ReadTimeout:     time.Second * 10,
			WriteTimeout:    time.Second * 10,
			ShutdownTimeout: time.Second * 10,
			TokenTTL:        15000,
			MainPath:        "/api/",
		},
		DB: DB{
			URL:          "localhost:5432",
			SchemaName:   "",
			MaxOpenConns: 2,
			MaxIdleConns: 2,
		},
	}
}

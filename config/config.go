package config

import (
	"errors"
	"log"
	"time"

	"github.com/ory/viper"
)

type Config struct {
	API API
	DB  DB
	JWT JWT
}

type API struct {
	Address         string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
	MainPath        string
	AccessTokenTTL  int64
	RefreshTokenTTL int64
	JwtKey          string
}

type JWT struct {
	AccessSecret  string
	RefreshSecret string
}

type DB struct {
	URL          string
	SchemaName   string
	MaxOpenConns int
	MaxIdleConns int
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}

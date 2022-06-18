package service

import (
	"errors"
	"fmt"

	"github.com/Mukhash/medods_auth/config"
	"github.com/Mukhash/medods_auth/internal/models"
	"github.com/Mukhash/medods_auth/internal/repository"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

type authService struct {
	logger *zap.Logger
	cfg    *config.Config
	repo   repository.Repository
}

func New(repo repository.Repository, cfg *config.Config, logger *zap.Logger) *authService {
	return &authService{repo: repo, cfg: cfg, logger: logger}
}

func (a *authService) Auth(payload string) (*models.Token, error) {
	user := &models.User{}
	var err error

	aToken := jwt.NewWithClaims(jwt.SigningMethodHS512, models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Local().Unix() + a.cfg.API.TokenTTL,
			IssuedAt:  jwt.TimeFunc().Local().Unix(),
		},
		UUID: payload,
	})
	rToken := jwt.NewWithClaims(jwt.SigningMethodHS512, models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Unix() + a.cfg.API.RefreshTokenTTL,
			IssuedAt:  jwt.TimeFunc().Local().Unix(),
		},
		UUID: payload,
	})

	tokens := &models.Token{}

	tokens.Access, err = aToken.SignedString([]byte(a.cfg.JWT.AccessSecret))
	if err != nil {
		return nil, err
	}

	tokens.Refresh, err = rToken.SignedString([]byte(a.cfg.JWT.AccessSecret))
	if err != nil {
		return nil, err
	}

	user.RefreshToken = tokens.Refresh
	user.UUID = payload

	if err = a.repo.InsertRefresh(user); err != nil {
		return nil, err
	}

	return tokens, nil
}

func (a *authService) Refresh(refreshToken string) (string, error) {
	claims, err := parseToken(refreshToken, []byte(a.cfg.JWT.RefreshSecret))
	if err != nil {
		return "", err
	}

	aToken := jwt.NewWithClaims(jwt.SigningMethodHS512, models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Local().Unix() + a.cfg.API.TokenTTL,
			IssuedAt:  jwt.TimeFunc().Local().Unix(),
		},
		UUID: claims.UUID,
	})

	return aToken.Raw, nil
}

func parseToken(mtoken string, signingKey []byte) (*models.Claims, error) {
	token, err := jwt.ParseWithClaims(mtoken, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		claims := &models.Claims{}
		return claims, err
	}

	claims, ok := token.Claims.(*models.Claims)
	if ok && token.Valid {
		return claims, nil
	}

	return claims, errors.New("invalid token...")
}

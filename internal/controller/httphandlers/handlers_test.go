package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/Mukhash/medods_auth/config"
	"github.com/Mukhash/medods_auth/internal/models"
	"github.com/Mukhash/medods_auth/internal/service/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"gopkg.in/go-playground/assert.v1"
)

func TestAuth(t *testing.T) {
	type data struct {
		UUID string `json:"guid"`
	}
	type Test struct {
		name                 string
		inputBody            data
		mockBehavior         func(r *mock.MockAuthService, client data)
		expectedStatusCode   int
		expectedResponseBody string
	}

	tests := [...]Test{
		{
			name:      "All is correct",
			inputBody: data{UUID: "123456789"},
			mockBehavior: func(r *mock.MockAuthService, client data) {
				r.EXPECT().CreateSession(client.UUID).Return(&models.Token{}, nil)
			},
			expectedStatusCode: 201,
		},
		{
			name:      "Empty guid",
			inputBody: data{UUID: ""},
			mockBehavior: func(r *mock.MockAuthService, client data) {
			},
			expectedStatusCode: 400,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			cfg := &config.Config{}
			logger := zap.NewExample()
			srv := mock.NewMockAuthService(c)
			handler := New(cfg, logger, srv)
			test.mockBehavior(srv, test.inputBody)

			e := echo.New()
			passRoutes := e.Group("/api")
			{
				passRoutes.POST("/auth", handler.Auth)
			}

			w := httptest.NewRecorder()
			body, _ := json.Marshal(test.inputBody)
			req := httptest.NewRequest("POST", "/api/auth", bytes.NewBufferString(string(body)))

			e.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatusCode, w.Code)

		})
	}
}

func TestRefresh(t *testing.T) {
	type data struct {
		RefreshToken string `json:"refresh_token"`
	}
	type Test struct {
		name                 string
		inputBody            data
		mockBehavior         func(r *mock.MockAuthService, client data)
		expectedStatusCode   int
		expectedResponseBody string
	}

	tests := [...]Test{
		{
			name:      "All is correct",
			inputBody: data{RefreshToken: "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTgyNDg4NTIsImlhdCI6MTY1NTY1Njg1MiwiVVVJRCI6IjEyMzQ1In0.HwHKoHMtwNTOzbZ-ps8Jhw7-8srjevEw_Oh6TKXVm97v2i65DvM-yKhmb9owmMfedkke6oMrd4EkI_PaG-stfw"},
			mockBehavior: func(r *mock.MockAuthService, client data) {
				r.EXPECT().Refresh(client.RefreshToken).Return("", nil)
			},
			expectedStatusCode: 201,
		},
		{
			name:      "Refresh broken",
			inputBody: data{RefreshToken: ""},
			mockBehavior: func(r *mock.MockAuthService, client data) {
				r.EXPECT().Refresh(client.RefreshToken).Return("", errors.New("invalid token..."))
			},
			expectedStatusCode: 500,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			cfg := &config.Config{}
			logger := zap.NewExample()
			srv := mock.NewMockAuthService(c)
			handler := New(cfg, logger, srv)
			test.mockBehavior(srv, test.inputBody)

			e := echo.New()
			passRoutes := e.Group("/api")
			{
				passRoutes.POST("/refresh", handler.Refresh)
			}

			w := httptest.NewRecorder()
			body, _ := json.Marshal(test.inputBody)
			req := httptest.NewRequest("POST", "/api/refresh", bytes.NewBufferString(string(body)))

			e.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatusCode, w.Code)

		})
	}
}

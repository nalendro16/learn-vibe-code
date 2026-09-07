package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nalendro16/learn-vibe-code/internal/dto"
	"github.com/nalendro16/learn-vibe-code/internal/service"
)

type mockUserService struct {
	registerFunc func(ctx context.Context, req dto.RegisterRequest) error
}

func (m *mockUserService) Register(ctx context.Context, req dto.RegisterRequest) error {
	if m.registerFunc != nil {
		return m.registerFunc(ctx, req)
	}
	return nil
}

func setupTestRouter(userService service.UserService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	handler := NewUserHandler(userService)
	router.POST("/api/register", handler.Register)
	return router
}

func TestUserHandler_Register(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		mockFunc       func(ctx context.Context, req dto.RegisterRequest) error
		expectedStatus int
		expectedResult string
	}{
		{
			name: "Success Registration",
			requestBody: map[string]string{
				"name":     "darsam",
				"email":    "darsam@gmail.com",
				"password": "password123",
			},
			mockFunc: func(ctx context.Context, req dto.RegisterRequest) error {
				return nil
			},
			expectedStatus: http.StatusCreated,
			expectedResult: "ok",
		},
		{
			name: "Validation Error - Missing Name",
			requestBody: map[string]string{
				"email":    "darsam@gmail.com",
				"password": "password123",
			},
			mockFunc:       nil,
			expectedStatus: http.StatusBadRequest,
			expectedResult: "error",
		},
		{
			name: "Validation Error - Invalid Email",
			requestBody: map[string]string{
				"name":     "darsam",
				"email":    "invalid-email",
				"password": "password123",
			},
			mockFunc:       nil,
			expectedStatus: http.StatusBadRequest,
			expectedResult: "error",
		},
		{
			name: "Validation Error - Password Too Short",
			requestBody: map[string]string{
				"name":     "darsam",
				"email":    "darsam@gmail.com",
				"password": "123",
			},
			mockFunc:       nil,
			expectedStatus: http.StatusBadRequest,
			expectedResult: "error",
		},
		{
			name: "Conflict - Email Already Registered",
			requestBody: map[string]string{
				"name":     "darsam",
				"email":    "darsam@gmail.com",
				"password": "password123",
			},
			mockFunc: func(ctx context.Context, req dto.RegisterRequest) error {
				return service.ErrEmailAlreadyRegistered
			},
			expectedStatus: http.StatusConflict,
			expectedResult: "error",
		},
		{
			name: "Internal Server Error",
			requestBody: map[string]string{
				"name":     "darsam",
				"email":    "darsam@gmail.com",
				"password": "password123",
			},
			mockFunc: func(ctx context.Context, req dto.RegisterRequest) error {
				return errors.New("database failure")
			},
			expectedStatus: http.StatusInternalServerError,
			expectedResult: "error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockSvc := &mockUserService{registerFunc: tc.mockFunc}
			router := setupTestRouter(mockSvc)

			bodyBytes, _ := json.Marshal(tc.requestBody)
			req, err := http.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(bodyBytes))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tc.expectedStatus {
				t.Errorf("expected status %d, got %d. Body: %s", tc.expectedStatus, w.Code, w.Body.String())
			}

			var resp dto.APIResponse
			if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
				t.Fatalf("failed to unmarshal response: %v", err)
			}

			if resp.Result != tc.expectedResult {
				t.Errorf("expected result '%s', got '%s'", tc.expectedResult, resp.Result)
			}
		})
	}
}

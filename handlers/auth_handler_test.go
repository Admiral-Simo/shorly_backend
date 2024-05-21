package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Admiral-Simo/shortly_backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserStore struct {
	mock.Mock
}

func (m *MockUserStore) CheckUser(username string, password string) (*models.User, error) {
	args := m.Called(username, password)
	if user := args.Get(0); user != nil {
		return user.(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserStore) CreateUser(username string, password string) (*models.User, error) {
	args := m.Called(username, password)
	if user := args.Get(0); user != nil {
		return user.(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserStore) GetUserById(id int) (*models.User, error) {
	args := m.Called(id)
	if user := args.Get(0); user != nil {
		return user.(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestAuthHandler_Login(t *testing.T) {
	app := fiber.New()
	mockUserStore := new(MockUserStore)
	authHandler := NewAuthHandler(mockUserStore)

	app.Post("/login", authHandler.Login)

	t.Run("Successful Login", func(t *testing.T) {
		mockUser := &models.User{Username: "testuser"}
		mockUserStore.On("CheckUser", "testuser", "password123").Return(mockUser, nil)

		loginRequest := AuthRequest{Username: "testuser", Password: "password123"}
		body, _ := json.Marshal(loginRequest)

		req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&response)

		assert.NotNil(t, response["user"])
		assert.Equal(t, "testuser", response["user"].(map[string]interface{})["username"])

		mockUserStore.AssertExpectations(t)
	})

	t.Run("Failed Login", func(t *testing.T) {
		mockUserStore.On("CheckUser", "testuser", "wrongpassword").Return(nil, errors.New("invalid password"))

		loginRequest := AuthRequest{Username: "testuser", Password: "wrongpassword"}
		body, _ := json.Marshal(loginRequest)

		req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		mockUserStore.AssertExpectations(t)
	})
}

func TestAuthHandler_Signup(t *testing.T) {
	app := fiber.New()
	mockUserStore := new(MockUserStore)
	authHandler := NewAuthHandler(mockUserStore)

	app.Post("/signup", authHandler.Signup)

	t.Run("Successful Signup", func(t *testing.T) {
		mockUser := &models.User{Username: "newuser"}
		mockUserStore.On("CreateUser", "newuser", "password123").Return(mockUser, nil)

		signupRequest := AuthRequest{Username: "newuser", Password: "password123"}
		body, _ := json.Marshal(signupRequest)

		req := httptest.NewRequest("POST", "/signup", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&response)

		assert.NotNil(t, response["user"])
		assert.Equal(t, "newuser", response["user"].(map[string]interface{})["username"])

		mockUserStore.AssertExpectations(t)
	})

	t.Run("Failed Signup", func(t *testing.T) {
		mockUserStore.On("CreateUser", "existinguser", "password123").Return(nil, errors.New("username already taken"))

		signupRequest := AuthRequest{Username: "existinguser", Password: "password123"}
		body, _ := json.Marshal(signupRequest)

		req := httptest.NewRequest("POST", "/signup", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		mockUserStore.AssertExpectations(t)
	})
}

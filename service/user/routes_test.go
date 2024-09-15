package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aormcuw/ecom/types"
	"github.com/gin-gonic/gin"
)

func TestUserServices(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if email is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "John",
			Lastname:  "Doe",
			Email:     "emailmail", // Invalid email
			Password:  "password123",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r := gin.Default()

		r.POST("/register", handler.HandleRegister)
		r.ServeHTTP(rr, req)

		// Expecting 400 response code for invalid email
		if rr.Code != 400 {
			t.Errorf("expected 400 response, got %d", rr.Code)
		}
	})
	t.Run("should correctly register the user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "John",
			Lastname:  "Doe",
			Email:     "john.doe@example.com",
			Password:  "password123",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r := gin.Default()

		r.POST("/register", handler.HandleRegister)
		r.ServeHTTP(rr, req)

		// Expecting 200 response code for invalid email
		if rr.Code != 200 {
			t.Errorf("expected 200 response, got %d", rr.Code)
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user %s not found", email)
}

func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}

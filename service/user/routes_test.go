package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aormcuw/ecom/types"
	"github.com/gin-gonic/gin"
)

func TestUserServices(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if user is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "John",
			Lastname:  "Doe",
			Email:     "example@example.com",
			Password:  "password123",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/api/v1/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		r := gin.Default()

		r.POST("/api/v1/register", handler.HandleRegister)
		r.ServeHTTP(rr, req)

		if rr.Code != 200 {
			t.Errorf("expected 200 response, got %d", rr.Code)
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}

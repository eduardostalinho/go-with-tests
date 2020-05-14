package httphandlers

import (
	"bytes"
	"io"
	"encoding/json"
	"reflect"
	"net/http"
	"net/http/httptest"
	"testing"
)
type MockUserService struct {
	RegisterFunc func(user User) (string, error)
	UsersRegistered []User
}
func (s *MockUserService) Register(user User) (string, error) {
	s.UsersRegistered = append(s.UsersRegistered, user)
	return s.RegisterFunc(user)
}

func TestRegisterUser(t *testing.T) {
	t.Run("register valid user", func (t *testing.T) {
		user := User{Name: "Erickson"}
		expectedID := "any"

		service := &MockUserService{
			RegisterFunc: func (user User) (string, error) {
				return expectedID, nil
			},
		}

		server := NewUserServer(service)

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", userToJSON(user))

		server.RegisterUser(res, req)

		if res.Body.String() != expectedID {
			t.Errorf("expected response to be %s, got %s", expectedID, res.Body.String())
		}

		if !reflect.DeepEqual(service.UsersRegistered[0], user) {
			t.Errorf("expected to register user %s, users registered: %v", user, service.UsersRegistered)
		}
	})
}

func userToJSON(user User) io.Reader {
	data, _ := json.Marshal(user)
	return bytes.NewReader(data)
}

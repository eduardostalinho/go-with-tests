package httphandlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type MockUserService struct {
	RegisterFunc    func(user User) (string, error)
	UsersRegistered []User
}

func (s *MockUserService) Register(user User) (string, error) {
	s.UsersRegistered = append(s.UsersRegistered, user)
	return s.RegisterFunc(user)
}

func TestRegisterUser(t *testing.T) {
	t.Run("register valid user", func(t *testing.T) {
		user := User{Name: "Erickson"}
		expectedID := "any"

		service := &MockUserService{
			RegisterFunc: func(user User) (string, error) {
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
	t.Run("cannot register invalid data", func(t *testing.T) {
		data := bytes.NewReader([]byte("blabla"))

		server := NewUserServer(nil)

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", data)

		server.RegisterUser(res, req)
		assertStatus(t, res, http.StatusBadRequest)
	})
}

func userToJSON(user User) io.Reader {
	data, _ := json.Marshal(user)
	return bytes.NewReader(data)
}

func assertStatus(t *testing.T, response *httptest.ResponseRecorder, want int) {
	t.Helper()
	got := response.Code
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

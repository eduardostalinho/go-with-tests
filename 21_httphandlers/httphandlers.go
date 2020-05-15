package httphandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string
}

type UserService interface {
	Register(user User) (id string, err error)
}

type UserServer struct {
	service UserService
}

func NewUserServer(s UserService) *UserServer {
	return &UserServer{service: s}
}

func (s *UserServer) RegisterUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, fmt.Sprintf("could not decode user payload, %v", err), http.StatusBadRequest)
		return
	}

	userID, err := s.service.Register(newUser)

	if err != nil {
		http.Error(w, fmt.Sprintf("unable to register user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, userID)
}

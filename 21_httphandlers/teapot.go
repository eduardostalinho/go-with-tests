package httphandlers

import (
	"net/http"
)

func Teapot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}
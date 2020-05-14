package httphandlers

import (
	"net/http"
	"testing"
	"net/http/httptest"
)


func TestTeapotHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	Teapot(res, req)

	if res.Code != http.StatusTeapot {
		t.Errorf("expected status %d, got %d", http.StatusTeapot, res.Code)
	}

}
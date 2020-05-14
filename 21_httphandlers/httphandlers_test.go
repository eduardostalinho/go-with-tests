package httphandlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTeapotHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	Teapot(res, req)

	if res.Code != http.StatusTeapot {
		t.Errorf("expected status %d, got %d", http.StatusTeapot, res.Code)
	}

}

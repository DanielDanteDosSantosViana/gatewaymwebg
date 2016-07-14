package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"tile_auth/handlers"
)

func TestRequestProxy(t *testing.T) {
	proxy := handlers.NewProxy()
	req, _ := http.NewRequest("POST", "", nil)
	w := httptest.NewRecorder()
	proxy.ProxyRequest(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Proxy Request didn't return %v", w.Code)
	}
}

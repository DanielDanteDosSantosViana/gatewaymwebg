package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestProxy(t *testing.T) {
	proxy := NewProxy()
	req, _ := http.NewRequest("GET", "/exemplo1", nil)
	w := httptest.NewRecorder()
	proxy.ProxyRequest(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Proxy Request didn't return %v", w.Code)
	}
}

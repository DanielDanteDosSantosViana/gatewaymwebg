package test

import (
	"github.com/DanielDanteDosSantosViana/gatewaymwebg/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestProxy(t *testing.T) {
	proxy := handlers.NewProxy()
	req, _ := http.NewRequest("GET", "/exemplo1", nil)
	w := httptest.NewRecorder()
	proxy.ProxyRequest(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Proxy Request didn't return. Status Code: %v", w.Code)
	}
}

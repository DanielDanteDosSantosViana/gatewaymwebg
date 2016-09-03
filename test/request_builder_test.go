package test

import (
	"github.com/DanielDanteDosSantosViana/gatewaymwebg/build"
	"net/http"
	"testing"
)

func TestRequestBuilder(t *testing.T) {
	req, _ := http.NewRequest("GET", "/exemplo1", nil)
	builder := build.NewRequestBuilder(req)
	_, err := builder.Build()
	if err != nil {
		t.Errorf("Error request return is %v ", err)
	}
}

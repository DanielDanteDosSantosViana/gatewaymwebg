package handlers

import (
	"github.com/DanielDanteDosSantosViana/gatewaymwebg/build"
	"net/http"
)

type Proxy struct {
}

func NewProxy() *Proxy {
	return &Proxy{}
}

func (p *Proxy) ProxyRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

	builder := build.NewRequestBuilder(r)
	builder.Build()
}

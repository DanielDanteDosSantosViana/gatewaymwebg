package handlers

import (
	"fmt"
	"github.com/DanielDanteDosSantosViana/gatewaymwebg/build"
	"github.com/DanielDanteDosSantosViana/gatewaymwebg/util"
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
	req, err := builder.Build()

	if err != nil {
		fmt.Errorf("error proxy request", err)
	}
	sender := util.NewSenderHTTP()
	_, err = sender.Send(req)
	if err != nil {
		fmt.Errorf("Error to send http", err)
	}
}

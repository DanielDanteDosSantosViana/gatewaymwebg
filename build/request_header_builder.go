package build

import (
	"github.com/DanielDanteDosSantosViana/gatewaymwebg/config"
	"net/http"
)

type RequestHeaderBuilder struct {
	next   RequestBuild
	Config config.Config
}

func (rh *RequestHeaderBuilder) Build(rq *http.Request) *http.Request {
	rq.Header.Set("Content-Type", "application/json")
	return rh.next.Build(rq)
}
func (rh *RequestHeaderBuilder) Previous(previous RequestBuild) {
	rh.next = previous
}
func NewHeaderBuilder(conf config.Config) *RequestHeaderBuilder {
	return &RequestHeaderBuilder{nil, conf}
}

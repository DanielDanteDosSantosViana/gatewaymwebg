package build

import (
	"github.com/DanielDanteDosSantosViana/gatewaymwebg/config"
	"net/http"
)

type RequestMethodBuilder struct {
	next        RequestBuild
	serviceConf config.Service
}

func (rm *RequestMethodBuilder) Build(rq *http.Request) (*http.Request, error) {
	rq.Method = rm.serviceConf.Method
	return rm.next.Build(rq)
}

func (rm *RequestMethodBuilder) Previous(previous RequestBuild) {
	rm.next = previous
}
func NewMethodBuilder(serviceConf config.Service) *RequestMethodBuilder {
	return &RequestMethodBuilder{nil, serviceConf}
}

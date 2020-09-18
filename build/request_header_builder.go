package build

import (
	"errors"
	"gatewaymwebg/config"
	"net/http"
)

var (
	ErrCreateHeaderToRequest = errors.New("Error to create Header in request")
)

type RequestHeaderBuilder struct {
	next        RequestBuild
	serviceConf config.Service
}

func (rh *RequestHeaderBuilder) Build(rq *http.Request) (*http.Request, error) {

	if rq == nil {
		return nil, ErrCreateHeaderToRequest
	}

	rq.Header = make(http.Header)
	rq.Header.Set("Content-Type", rh.serviceConf.Type)

	return rh.next.Build(rq)
}
func (rh *RequestHeaderBuilder) Previous(previous RequestBuild) {
	rh.next = previous
}
func NewHeaderBuilder(serviceConf config.Service) *RequestHeaderBuilder {
	return &RequestHeaderBuilder{nil, serviceConf}
}

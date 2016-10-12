package build

import (
	"errors"
	"github.com/DanielDanteDosSantosViana/gatewaymwebg/config"
	"net/http"
)

var (
	ErrCreateParamToRequest = errors.New("Error create Param in Request Body")
)

type RequestParamBuilder struct {
	next        RequestBuild
	serviceConf config.Service
	rqReceiver  *http.Request
}

func (rp *RequestParamBuilder) Build(rq *http.Request) (*http.Request, error) {
	if rq == nil {
		return nil, ErrCreateParamToRequest
	}
	if rp.rqReceiver.Body == nil {
		return nil, ErrCreateParamToRequest
	} else {
		rq.Body = rp.rqReceiver.Body
		rq.ContentLength = rp.rqReceiver.ContentLength
	}
	return rq, nil
}
func (rp *RequestParamBuilder) Previous(previous RequestBuild) {
	rp.next = previous
}
func NewParamBuilder(serviceConf config.Service, rqReceiver *http.Request) *RequestParamBuilder {
	return &RequestParamBuilder{nil, serviceConf, rqReceiver}
}

package build

import (
	"errors"
	"github.com/DanielDanteDosSantosViana/gatewaymwebg/config"
	"net/http"
	"net/url"
)

var (
	ErrInvalidURL = errors.New("Invalid url")
)

type RequestUrlBuilder struct {
	next        RequestBuild
	serviceConf config.Service
}

func (rUrl *RequestUrlBuilder) Build(rq *http.Request) (*http.Request, error) {
	u, err := url.Parse(rUrl.serviceConf.UrlDest)
	if err != nil {
		return nil, ErrInvalidURL
	}
	rq.URL = u
	return rq, nil

}

func (rUrl *RequestUrlBuilder) Previous(previous RequestBuild) {
	rUrl.next = previous
}

func NewUrlBuilder(serviceConf config.Service) *RequestUrlBuilder {
	return &RequestUrlBuilder{nil, serviceConf}
}

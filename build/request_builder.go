package build

import (
	"errors"
	"github.com/DanielDanteDosSantosViana/gatewaymwebg/config"
	"net/http"
)

var (
	ErrCreateRequestBuilder = errors.New("Request Builder error in make request")
)

type RequestBuilder struct {
	req *http.Request
}

func NewRequestBuilder(req *http.Request) *RequestBuilder {
	rb := new(RequestBuilder)
	rb.req = req
	return rb
}

func (rb *RequestBuilder) Build() (*http.Request, error) {
	service, err := config.FindServiceIn(rb.req)
	if err != nil {
		return nil, err
	}

	rUrl := NewUrlBuilder(service)
	rh := NewHeaderBuilder(service)
	rUrl.Previous(rh)

	return rUrl.Build(new(http.Request))
}

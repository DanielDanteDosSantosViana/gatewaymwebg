package build

import (
	"errors"
	"gatewaymwebg/config"
	"net/http"
)

var (
	ErrCreateRequestBuilder = errors.New("Request Builder error in make request")
)

type RequestBuilder struct {
	req *http.Request
}

func NewRequestBuilder(request *http.Request) *RequestBuilder {
	rb := new(RequestBuilder)
	rb.req = request
	return rb
}

func (rb *RequestBuilder) Build() (*http.Request, error) {
	service, err := config.FindServiceIn(rb.req)
	if err != nil {
		return nil, err
	}

	rUrl := NewUrlBuilder(service)
	rm := NewMethodBuilder(service)
	rh := NewHeaderBuilder(service)
	pb := NewParamBuilder(service, rb.req)

	rUrl.Previous(rm)
	rm.Previous(rh)
	rh.Previous(pb)
	return rUrl.Build(new(http.Request))
}

package build

import (
	"fmt"
	"github.com/DanielDanteDosSantosViana/gatewaymwebg/config"
	"net/http"
)

type RequestBuilder struct {
	req *http.Request
}

func NewRequestBuilder(req *http.Request) *RequestBuilder {
	rb := new(RequestBuilder)
	rb.req = req
	return rb
}

func (rb *RequestBuilder) Build() *http.Request {
	configuration := config.New()
	conf, err := configuration.Load()
	if err != nil {
		fmt.Print(err)
	}
	rUrl := NewUrlBuilder(conf)
	rh := NewHeaderBuilder(conf)

	rUrl.Previous(rh)

	return rUrl.Build(new(http.Request))
}

func (rb *RequestBuilder) defineChain() {

}

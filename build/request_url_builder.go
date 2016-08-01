package build

import (
	"fmt"
	"github.com/DanielDanteDosSantosViana/gatewaymwebg/config"
	"net/http"
	"net/url"
)

type RequestUrlBuilder struct {
	next   RequestBuild
	Config config.Config
}

func (rUrl *RequestUrlBuilder) Build(rq *http.Request) *http.Request {
	u, err := url.Parse("http://bing.com/search?q=dotnet")
	rq.URL = u
	fmt.Errorf("Could not decode config.toml: %v", err)
	return rq
}
func (rUrl *RequestUrlBuilder) Previous(previous RequestBuild) {
	rUrl.next = previous
}

func NewUrlBuilder(conf config.Config) *RequestUrlBuilder {
	return &RequestUrlBuilder{nil, conf}
}

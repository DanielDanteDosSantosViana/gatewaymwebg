package build

import "net/http"

type RequestBuild interface {
	Build(rq *http.Request) *http.Request
	Previous(previous RequestBuild)
}

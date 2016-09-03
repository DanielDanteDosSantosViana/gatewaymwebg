package handlers

import (
	"github.com/DanielDanteDosSantosViana/gatewaymwebg/build"
	"github.com/DanielDanteDosSantosViana/gatewaymwebg/config"
	sender "github.com/DanielDanteDosSantosViana/gatewaymwebg/util"
	"io/ioutil"
	"log"
	"net/http"
)

type ProxyHandler struct {
}

func NewProxy() *ProxyHandler {
	return &ProxyHandler{}
}

func (ph *ProxyHandler) ProxyRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

	if isFavicoRequest(r) {
		w.Header().Set("Content-Type", "image/x-icon'")
		w.WriteHeader(http.StatusOK)
		return
	}

	builder := build.NewRequestBuilder(r)
	req, err := builder.Build()

	if err != nil {
		if err == config.ErrorServiceNotFound {
			logError("config file", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		logError("internal Server", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	resp, err := sender.Send(req)
	if err != nil {
		logError("when send http", err)
		w.WriteHeader(resp.StatusCode)
		return
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	w.WriteHeader(resp.StatusCode)
	w.Write(result)
}

func isFavicoRequest(r *http.Request) bool {
	if r.URL.Path == "/favicon.ico" {
		return true
	}
	return false
}

func logError(msg string, err error) {
	log.Printf("Error: %s : %v", msg, err)
}

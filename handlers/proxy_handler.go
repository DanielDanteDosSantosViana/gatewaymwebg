package handlers

import (
	"fmt"
	"gatewaymwebg/build"
	"gatewaymwebg/config"
	sender "gatewaymwebg/util"
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
			fmt.Fprintf(w, "Config file Error: %v", err)
			return
		}
		logError("internal Server", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal Server : %v", err)
		return
	}

	resp, errResp := sender.Send(req)
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logError("Error when read response", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error when read response : %v", err)
		return
	}

	if errResp != nil {
		logError("when send http", err)
		w.WriteHeader(resp.StatusCode)
		w.Write(result)
		return
	}

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

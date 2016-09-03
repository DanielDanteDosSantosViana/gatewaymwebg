package config

import (
	"errors"
	"log"
	"net/http"
)

var (
	ErrorServiceNotFound = errors.New("Service Not Found")
)

func FindServiceIn(req *http.Request) (Service, error) {
	loader := &ConfigurationLoader{}
	conf, error := loader.Load()
	if error != nil {
		return Service{}, error
	}

	url := req.URL.String()
	log.Printf("URL do request : %v", url)
	method := req.Method
	services := conf.Services
	for _, v := range services {
		if v.UrlSrc == url {
			if v.Method == method {
				return v, nil
			}
		}
	}
	return Service{}, ErrorServiceNotFound
}

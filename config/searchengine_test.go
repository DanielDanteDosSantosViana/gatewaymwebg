package config

import (
	"net/http"
	"testing"
)

func TestShouldFindServiceInConfigFile(t *testing.T) {
	req, _ := http.NewRequest("GET", "/exemplo1", nil)
	service, err := FindServiceIn(req)
	if (Service{}) == service {
		t.Errorf("Error not found the service, but the request parameters exist in the configuration file :  %v", err)
	}

}

func TestShouldThrowsErrorWithInvalidParameterInRequest(t *testing.T) {
	req, _ := http.NewRequest("POST", "/exemplo1", nil)
	_, err := FindServiceIn(req)
	if err != ErrorServiceNotFound {
		t.Errorf("Because error did not show the error 'Service Not Found'")
	}

}

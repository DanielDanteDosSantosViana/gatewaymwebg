package util

import (
	"errors"
	"log"
	"net/http"
	"strings"
)

var (
	ErrorContentTypeNotFound = errors.New("Contet-Type Not Supported in this application")
)

func Send(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error send %v", err)
		return nil, errors.New("Error sending request to the url : " + req.URL.String())
	}
	isRespOk, err := isValidResponse(resp)
	if isRespOk {
		return resp, nil
	} else {
		return resp, err
	}
}

func isJson(str string) bool {
	return strings.Contains(str, "application/json")

}

func isValidResponse(resp *http.Response) (bool, error) {
	tipo := resp.Header.Get("Content-Type")
	if isJson(tipo) {
		if resp.StatusCode != http.StatusOK {
			return false, errors.New("Status code is not OK : Status Code : " + resp.Status)
		} else {
			return true, nil
		}
	}
	return false, ErrorContentTypeNotFound
}

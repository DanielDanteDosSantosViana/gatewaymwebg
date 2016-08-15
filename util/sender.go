package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	SenderHTTP struct{}
)

func NewSenderHTTP() *SenderHTTP {
	return &SenderHTTP{}
}

func (send *SenderHTTP) Send(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return resp, errors.New("Error sending request to the url : " + req.URL.String())
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return resp, nil
}

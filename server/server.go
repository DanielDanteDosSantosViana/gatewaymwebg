package server

import (
	"log"
	"net/http"
	"tile_auth/handlers"
)

const public = "server_cert.pem"
const privkey = "server_key.pem"

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Listen(url string) {
	init()
	err := http.ListenAndServeTLS(url, public, privkey, nil)
	if err != nil {
		log.Fatal("Error init Server : ", err)
	}
}

func init() {
	proxy := handlers.NewProxy()
	http.HandleFunc("/", http.HandlerFunc(proxy.ProxyRequest))
}

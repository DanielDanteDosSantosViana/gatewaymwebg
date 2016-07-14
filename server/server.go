package server

import (
	"github.com/DanielDanteDosSantosViana/tile_auth/handlers"
	"log"
	"net/http"
)

const public = "server_cert.pem"
const privkey = "server_key.pem"

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Listen(url string) {
	s.init()
	err := http.ListenAndServeTLS(url, public, privkey, nil)
	if err != nil {
		log.Fatal("Error init Server : ", err)
	}
}

func (s *Server) init() {
	proxy := handlers.NewProxy()
	http.HandleFunc("/", http.HandlerFunc(proxy.ProxyRequest))
}

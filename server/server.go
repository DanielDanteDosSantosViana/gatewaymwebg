package server

import (
	"gatewaymwebg/handlers"
	"log"
	"net/http"
)


type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Listen(url string) {
	s.init()
	err := http.ListenAndServe(url,nil)
	if err != nil {
		log.Fatal("Error init Server : ", err)
	}
}

func (s *Server) init() {
	proxy := handlers.NewProxy()
	http.HandleFunc("/", proxy.ProxyRequest)
}

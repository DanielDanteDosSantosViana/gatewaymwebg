package main

import (
	"github.com/DanielDanteDosSantosViana/gatewaymwebg/server"
)

func main() {
	server := server.NewServer()
	server.Listen(":8080/gateway")
}

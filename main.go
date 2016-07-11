package main

import (
	"tile_auth/server"
)

func main() {
	server := server.NewServer()
	server.Listen(":8080/gateway")
}

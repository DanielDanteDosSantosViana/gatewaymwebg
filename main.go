package main

import (
	"gatewaymwebg/server"
	"log"
)

func main() {
	server := server.NewServer()
	log.Print("Gatewaymwebg is alive...")
	server.Listen(":8081")
}

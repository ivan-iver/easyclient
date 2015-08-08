package main

import (
	"client/core"
	"client/net"
	"log"
)

func main() {
	var err error
	config := core.Config{}
	if err = config.Init(); err != nil {
		log.Printf("Init: %v", err)
	}
	log.Printf("Config: %+v", config)
	request := net.NewRequest(config)
	if err = request.Send(); err != nil {
		log.Printf("Error: %v", err)
		return
	}
	log.Printf("Response: %v", request.Response)
}

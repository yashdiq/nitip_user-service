package main

import (
	"log"

	"github.com/yashdiq/nitip_user-service/pkg/di"
)

func main() {
	server, err := di.Initialize()

	if err != nil {
		log.Fatalln("Error in initializing the api", err)
	}
	server.Start()
}
package main

import (
	"log"
	"user-service/pkg/di"
)

func main() {
	server, err := di.Initialize()

	if err != nil {
		log.Fatalln("Error in initializing the api", err)
	}
	server.Start()
}
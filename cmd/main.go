package main

import (
	"log"

	"github.com/alberto-debug/HttpHandler/cmd/api"
)

func main() {
	server := api.NEWAPISERVER(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

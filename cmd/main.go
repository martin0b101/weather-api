package main

import (
	"log"

	"github.com/weather-api/cmd/api"
)





func main() {
	server := api.NewApiServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
package main

import (
	"github.com/DulioCagg/interview/server"
	"log"
)

func main() {
	server := server.New()
	defer server.DB.Close()

	log.Fatal(server.Router.Run(":8080"))
}

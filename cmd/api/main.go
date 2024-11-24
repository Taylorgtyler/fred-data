package main

import (
	"fmt"
	"log"

	"github.com/taylorgtyler/fred-data/internal/server"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}

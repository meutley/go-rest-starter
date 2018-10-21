package main

import (
	"log"

	"./server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	server := server.Server{}
	server.Start()
}

package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"my-oauth-server/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	go server.InitSocket()
	server.InitQueue()
	sv := server.Init()

	log.Printf("Service starting on %s \n\n\n", sv.Addr)

	err = sv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}

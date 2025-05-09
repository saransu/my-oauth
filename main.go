package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"konsys.co/ks-go-http/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	sv := server.Init()
	log.Printf("Service starting on %s \n\n\n", sv.Addr)

	err = sv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}

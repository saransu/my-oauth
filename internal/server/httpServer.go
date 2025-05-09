package server

import (
	"fmt"
	"net/http"
	"os"
)

func Init() *http.Server {
	mux := newMux()

	loggedMux := logWrapper(mux)
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: loggedMux,
	}

	return &server
}

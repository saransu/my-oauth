package server

import (
	"fmt"
	"net/http"
	"os"
)

func Init() *http.Server {
	mux := newMux()

	loggedMux := logWrapper(mux)
	handler := defaultWrapper(loggedMux)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: handler,
	}

	return &server
}

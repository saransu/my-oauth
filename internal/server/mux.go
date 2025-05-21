package server

import (
	"my-oauth-server/internal/oauth"
	"net/http"
)

func newMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", healthCheck)

	oauth.RegisterOAuthController(mux)

	return mux
}

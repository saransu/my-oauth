package server

import (
	"my-oauth-server/internal/utils"
	"net/http"
)

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	body := utils.ResponseBody{
		Message: "Health Check OK",
	}

	w.WriteHeader(http.StatusOK)
	utils.WriteResponseBody(w, body)
	return
}

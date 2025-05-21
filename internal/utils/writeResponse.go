package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseBody struct {
	Data    any    `json:"data"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func WriteResponseBody(w http.ResponseWriter, body ResponseBody) {
	jsonResp, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "something went wrong")
	}

	_, err = w.Write(jsonResp)
	if err != nil {
		fmt.Fprintf(w, "something went wrong")
	}
}

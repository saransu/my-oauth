package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body := map[string]string{
		"message": "Health Check OK",
	}

	rBody, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(500)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Fprintf(w, "something went wrong")
		}
	}

	w.WriteHeader(200)
	_, err = w.Write(rBody)
	if err != nil {
		fmt.Fprintf(w, "something went wront")
	}
}

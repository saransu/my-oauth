package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	body := map[string]string{
		"message": "Health Check OK",
	}

	rBody, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Fprintf(w, "something went wrong")
		}
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(rBody)
	if err != nil {
		fmt.Fprintf(w, "something went wront")
	}
}

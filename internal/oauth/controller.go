package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterOAuthController(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/oauth", handleGrant)
	mux.HandleFunc("GET /api/v1/oauth/token", grantAccessToken)
}

func handleGrant(w http.ResponseWriter, req *http.Request) {
	rawGT := req.URL.Query().Get("grant_type")

	gt, err := parseGrantType(rawGT)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		r := map[string]string{
			"error": err.Error(),
		}
		rb, err := json.Marshal(r)
		if err != nil {
			fmt.Fprintf(w, "something went wrong")
		}

		w.Write(rb)
		return
	}

	switch gt {
	case authorizationCode:
		grantAuthorizationCode(w, req)
	default:
		break
	}
}

package oauth

import (
	"my-oauth-server/internal/utils"
	"net/http"
)

func RegisterController(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/oauth", HandleGrant)
	mux.HandleFunc("GET /api/v1/oauth/token", GrantAccessToken)
}

func HandleGrant(w http.ResponseWriter, req *http.Request) {
	rawGT := req.URL.Query().Get("grant_type")

	gt, err := parseGrantType(rawGT)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.WriteResponseBody(w, utils.ResponseBody{Data: nil, Error: err.Error()})
		return
	}

	switch gt {
	case authorizationCode:
		grantAuthorizationCode(w, req)
	default:
		break
	}
}

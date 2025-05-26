package oauth

import (
	"my-oauth-server/internal/utils"
	"net/http"
)

const clientID = "ZRPY48s9kjLb2Pr5IY5zCgYKxWk30g3V"
const clientSecret = "H1XuhfthXcTuQ0IKzH28hVSMsWgoIpju"

var authCode string

func grantAuthorizationCode(w http.ResponseWriter, req *http.Request) {
	cID := req.URL.Query().Get("client_id")

	if cID != clientID {
		w.WriteHeader(http.StatusUnauthorized)
		utils.WriteResponseBody(w, utils.ResponseBody{Error: "invalid client_id"})
		return
	}

	authCode = utils.NewRandomString(64)
	resp := map[string]string{
		"code": authCode,
	}
	utils.WriteResponseBody(w, utils.ResponseBody{Message: "success", Data: resp})
}

func grantAccessToken(w http.ResponseWriter, req *http.Request) {
	code := req.URL.Query().Get("code")
	cID := req.URL.Query().Get("client_id")
	cSecret := req.URL.Query().Get("client_secret")

	if code != authCode {
		w.WriteHeader(http.StatusUnauthorized)
		utils.WriteResponseBody(w, utils.ResponseBody{Error: "invalid code"})
		return
	}

	if cID != clientID {
		w.WriteHeader(http.StatusUnauthorized)
		utils.WriteResponseBody(w, utils.ResponseBody{Error: "invalid client_id"})
		return
	}

	if cSecret != clientSecret {
		w.WriteHeader(http.StatusUnauthorized)
		utils.WriteResponseBody(w, utils.ResponseBody{Error: "invalid client_secret"})
		return
	}

	u := utils.User{
		ID:   1,
		Name: "Saran",
		Age:  20,
	}

	token, err := utils.GenerateToken(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.WriteResponseBody(w, utils.ResponseBody{Error: err.Error()})
	}
	utils.WriteResponseBody(w, utils.ResponseBody{
		Data: token,
	})
}

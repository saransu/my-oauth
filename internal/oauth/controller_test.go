package oauth_test

import (
	"encoding/json"
	"fmt"
	"log"
	"my-oauth-server/internal/oauth"
	"my-oauth-server/internal/utils"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"runtime"
	"testing"
)

type testCase struct {
	description string
	qs          map[string]string
	statusCode  int
}

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..", "..")
	err := os.Chdir(dir)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestGetOauth(t *testing.T) {
	tests := []testCase{
		{
			description: "invalid grant type",
			qs: map[string]string{
				"grant_type": "invalid",
			},
			statusCode: http.StatusBadRequest,
		},
		{
			description: "invalid client_id",
			qs: map[string]string{
				"grant_type": "authorization_code",
				"client_id":  "invalid",
			},
			statusCode: http.StatusUnauthorized,
		},
		{
			description: "valid client_id",
			qs: map[string]string{
				"grant_type": "authorization_code",
				"client_id":  oauth.ClientID,
			},
			statusCode: http.StatusOK,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.description, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/api/v1/oauth", nil)
			q := req.URL.Query()
			for k, v := range test.qs {
				q.Set(k, v)
			}
			req.URL.RawQuery = q.Encode()

			w := httptest.NewRecorder()

			oauth.HandleGrant(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			if resp.StatusCode != test.statusCode {
				t.Errorf("expected %d but got %d", test.statusCode, resp.StatusCode)
			}
		})
	}
}

func TestGetOauthToken(t *testing.T) {
	r, _ := http.NewRequest(http.MethodGet, "/api/v1/oauth", nil)
	q := r.URL.Query()
	q.Set("grant_type", "authorization_code")
	q.Set("client_id", oauth.ClientID)

	r.URL.RawQuery = q.Encode()

	w := httptest.NewRecorder()
	oauth.HandleGrant(w, r)

	res := w.Result()
	defer res.Body.Close()

	var jsonResp utils.ResponseBody
	_ = json.NewDecoder(res.Body).Decode(&jsonResp)

	fmt.Println(jsonResp)

	authCode, ok := jsonResp.Data.(string)
	if !ok {
		t.Errorf("cannot retrieve authorization code")
		return
	}

	tests := []testCase{
		{
			description: "valid",
			qs: map[string]string{
				"client_id":     oauth.ClientID,
				"client_secret": oauth.ClientSecret,
				"code":          authCode,
			},
			statusCode: http.StatusOK,
		},
		{
			description: "invalid_auth_code",
			qs: map[string]string{
				"client_id":     oauth.ClientID,
				"client_secret": oauth.ClientSecret,
				"code":          "",
			},
			statusCode: http.StatusUnauthorized,
		},
		{
			description: "invalid_client_id",
			qs: map[string]string{
				"client_id":     "invalid",
				"client_secret": oauth.ClientSecret,
				"code":          authCode,
			},
			statusCode: http.StatusUnauthorized,
		},
		{
			description: "invalid_client_secret",
			qs: map[string]string{
				"client_id":     oauth.ClientID,
				"client_secret": "invalid",
				"code":          authCode,
			},
			statusCode: http.StatusUnauthorized,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.description, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/api/v1/oauth/token", nil)
			q := req.URL.Query()
			for k, v := range test.qs {
				q.Set(k, v)
			}

			req.URL.RawQuery = q.Encode()
			w := httptest.NewRecorder()

			oauth.GrantAccessToken(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			if resp.StatusCode != test.statusCode {
				t.Errorf("expected %d but got %d", test.statusCode, resp.StatusCode)
			}
		})
	}
}

package oauth

import "fmt"

type grantType int

const (
	authorizationCode grantType = iota
)

func parseGrantType(gt string) (grantType, error) {
	switch gt {
	case "authorization_code":
		return authorizationCode, nil
	default:
		return -1, fmt.Errorf("invalid grant type")
	}
}

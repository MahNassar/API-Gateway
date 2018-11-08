package auth

import (
	"net/http"
	"fmt"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"encoding/json"
	"encoding/base64"
)

func CheckAuth(r *http.Request, auth bool) (string, error) {
	if auth {
		token := getToken(r)
		if token == "" {
			err := fmt.Errorf("Token not found")
			return "", err
		}
		tokenData, err := parseToken(token)
		if err != nil {
			return "", err
		}
		base64encoded := base64.StdEncoding.EncodeToString([]byte(tokenData))
		encText := encryption(base64encoded)
		return encText, nil
	}
	return "", nil
}

func getToken(r *http.Request) string {
	var token string
	token = r.Header.Get("Authorization")
	return token
}

func parseToken(tokenString string) (string, error) {
	tokenList := strings.Split(tokenString, "Bearer")
	if len(tokenList) != 2 {
		err := fmt.Errorf("Token not found")
		return "", err
	}
	tokenString = strings.TrimSpace(tokenList[1])

	signingKey := "secrit"

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if err != nil {
		return "", err
	}
	var tokenData string
	for key, val := range claims {
		if key == "sub" {
			data, _ := json.Marshal(val)
			tokenData = string(data)
		}
	}
	return tokenData, nil
}

package jwt_token

import (
	"errors"

	"github.com/NunChatSpace/client-service/config"
	"github.com/golang-jwt/jwt"
)

type AccessTokenClaims struct {
	jwt.StandardClaims
	Permission []string `json:"permission"`
	Type       string   `json:"type"`
	UserID     string   `json:"user_id"`
}

func Decode(refreshToken string, config *config.Config) (AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Credentials.SignString), nil
	})
	if err != nil {
		return AccessTokenClaims{}, err
	}

	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok {
		return AccessTokenClaims{}, errors.New("jwt parsing error")
	}

	return *claims, nil
}

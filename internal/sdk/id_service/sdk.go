package idservice

import (
	"encoding/json"
	"net/http"

	"github.com/NunChatSpace/client-service/config"
	"github.com/NunChatSpace/client-service/internal/response_wrapper"
)

type IntrospectionData struct {
	Data    TokenData `json:"data"`
	Message string    `json:"message"`
}

type TokenData struct {
	Expire     string   `json:"exp"`
	Permission []string `json:"permission"`
	Type       string   `json:"type"`
	UserID     string   `json:"user_id"`
	VerifyCode string   `json:"verify_code"`
}

func HealthCheck(config *config.Config) response_wrapper.Model {
	endpoint := config.Services.IdentityService + "/_healthz"
	resp, err := http.Get(endpoint)
	if err != nil {
		return response_wrapper.Model{
			ErrorCode: http.StatusInternalServerError,
			Message:   err.Error(),
		}
	}

	if resp.StatusCode != http.StatusOK {
		return response_wrapper.Model{
			ErrorCode: resp.StatusCode,
			Message:   "Identtity Service unhealthy",
		}
	}

	return response_wrapper.Model{
		ErrorCode: resp.StatusCode,
		Message:   "Identtity Service healthy",
	}
}

func VerifyToken(config *config.Config, token string) TokenData {
	endpoint := config.Services.IdentityService + "/intrspct"
	client := &http.Client{}
	var result TokenData
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return result
	}

	req.Header.Set("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		return result
	}

	if resp.StatusCode != http.StatusOK {
		return result
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	defer resp.Body.Close()
	if err != nil {
		return result
	}

	return result
}

package idservice

import (
	"net/http"

	"github.com/NunChatSpace/client-service/config"
	"github.com/NunChatSpace/client-service/internal/response_wrapper"
)

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

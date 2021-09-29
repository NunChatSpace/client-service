package config

import (
	"context"

	"github.com/NunChatSpace/client-service/config"
	"github.com/gin-gonic/gin"
)

type configContext struct{}

func FromContext(c context.Context) config.Config {
	val := c.Value(&configContext{})
	if val == nil {
		return config.Config{}
	}

	return val.(config.Config)
}

func Handler(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), &configContext{}, config)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

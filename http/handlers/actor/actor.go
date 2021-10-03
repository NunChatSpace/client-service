package actor

import (
	"context"
	"net/http"

	"github.com/NunChatSpace/client-service/config"
	idservice "github.com/NunChatSpace/client-service/internal/sdk/id_service"
	"github.com/gin-gonic/gin"
)

type actorContext struct{}

func FromContext(c context.Context) Model {
	val := c.Value(&actorContext{})
	if val == nil {
		return Model{}
	}

	return val.(Model)
}

func Handler(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header["Authorization"]
		if len(tokenHeader) == 0 {
			c.String(http.StatusForbidden, "unauthenticate")
			c.AbortWithStatus(http.StatusForbidden)
			return
		} else {
			result := idservice.VerifyToken(config, tokenHeader[0])
			if result.UserID == "" {
				c.String(http.StatusForbidden, "unauthenticate")
				c.AbortWithStatus(http.StatusForbidden)
				return
			}
			actor := Model{
				UserID:     result.UserID,
				Permission: result.Permission,
			}

			ctx := context.WithValue(c.Request.Context(), &actorContext{}, actor)
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		}
	}
}

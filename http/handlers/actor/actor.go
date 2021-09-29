package actor

import (
	"context"
	"strings"

	"github.com/NunChatSpace/client-service/config"
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
			ctx := context.WithValue(c.Request.Context(), &actorContext{}, Model{})
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		} else {
			token := strings.Split(tokenHeader[0], " ")[1]
			actor := GetActorFromToken(token, config)
			ctx := context.WithValue(c.Request.Context(), &actorContext{}, actor)
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		}
	}
}

package handlers

import (
	"github.com/NunChatSpace/client-service/config"
	"github.com/NunChatSpace/client-service/http/handlers/actor"
	cfgHandler "github.com/NunChatSpace/client-service/http/handlers/config"
	"github.com/NunChatSpace/client-service/http/handlers/cors"
	"github.com/gin-gonic/gin"
)

func SetHandlers(config *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(cors.Handler())
	// router.Use(database.Handler(db))
	router.Use(cfgHandler.Handler(config))
	router.Use(actor.Handler(config))

	return router
}

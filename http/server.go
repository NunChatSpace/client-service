package http

import (
	"github.com/NunChatSpace/client-service/config"
	"github.com/NunChatSpace/client-service/http/handlers"
	"github.com/NunChatSpace/client-service/internal/routes/routes_users"
	idservice "github.com/NunChatSpace/client-service/internal/sdk/id_service"
	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
	// db, err := entities.NewDB()
	// if err != nil {
	// 	panic(err)
	// }

	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	router := handlers.SetHandlers(config)
	routes_users.AddGroup(router)

	router.GET("/_healthz", func(c *gin.Context) {
		id_result := idservice.HealthCheck(config)
		c.JSON(200, gin.H{
			"message": gin.H{
				"id_service": id_result.Message,
			},
		})
	})

	return router
}

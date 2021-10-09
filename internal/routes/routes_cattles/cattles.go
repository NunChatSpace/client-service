package routes_cattles

import (
	"encoding/json"
	"net/http"

	"github.com/NunChatSpace/client-service/http/handlers/database"
	dc "github.com/NunChatSpace/client-service/internal/deliveries/deliveries_cattles"
	"github.com/NunChatSpace/client-service/internal/response_wrapper"
	"github.com/gin-gonic/gin"
)

func AddGroup(r *gin.Engine) {
	cattlesGroup := r.Group("/cattles")
	cattlesGroup.POST("/", add)
}

func add(c *gin.Context) {
	var regisModel dc.CattlesRegisterModel

	err := json.NewDecoder(c.Request.Body).Decode(&regisModel)
	if err != nil {
		response_wrapper.Resp(http.StatusInternalServerError, "", err.Error(), c)
		return
	}

	db := database.FromContext(c.Request.Context())
	if db == nil {
		response_wrapper.Resp(http.StatusInternalServerError, "", "db does not exist", c)
		return
	}

	m, err := dc.AddCattles(db, regisModel)
	if err != nil {
		response_wrapper.Resp(m.ErrorCode, m.Data, err.Error(), c)
		return
	}
	response_wrapper.Resp(m.ErrorCode, m.Data, "", c)
}

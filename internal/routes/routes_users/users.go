package routes_users

import (
	"encoding/json"
	"net/http"

	"github.com/NunChatSpace/client-service/http/handlers/database"
	du "github.com/NunChatSpace/client-service/internal/deliveries/deliveries_users"
	"github.com/NunChatSpace/client-service/internal/response_wrapper"
	"github.com/gin-gonic/gin"
)

func AddGroup(r *gin.Engine) {
	usersGroup := r.Group("/users")
	usersGroup.POST("/", register)
}

func register(c *gin.Context) {
	var regisModel du.UserRegisterModel

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

	m, err := du.Register(db, regisModel)
	if err != nil {
		response_wrapper.Resp(m.ErrorCode, m.Data, err.Error(), c)
		return
	}
	response_wrapper.Resp(m.ErrorCode, m.Data, "", c)
}

package deliveries_users

import (
	"net/http"

	"github.com/NunChatSpace/client-service/internal/deliveries"
	"github.com/NunChatSpace/client-service/internal/entities"
	"github.com/NunChatSpace/client-service/internal/response_wrapper"
)

type UserRegisterModel struct {
	FirstName   string `json:"first_name"`
	MiddleName  string `json:"middle_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}

func Register(db entities.DB, model UserRegisterModel) (response_wrapper.Model, error) {
	rolename, err := db.RoleName().Get("user")
	if err != nil {
		return deliveries.InternalError(UserRegisterModel{}, err)
	}

	cinfo := entities.ContactModel{
		Email:       model.Email,
		PhoneNumber: model.PhoneNumber,
		Address:     model.Address,
	}
	contact, err := db.Contact().Add(cinfo)
	if err != nil {
		return deliveries.InternalError(UserRegisterModel{}, err)
	}

	uinfo := entities.UserModel{
		FirstName:  model.FirstName,
		MiddleName: model.MiddleName,
		LastName:   model.LastName,
		ContactID:  contact.ID,
		RoleNameID: rolename.ID,
	}

	_, err = db.User().Add(uinfo)
	if err != nil {
		return deliveries.InternalError(UserRegisterModel{}, err)
	}

	return response_wrapper.Model{
		ErrorCode: http.StatusOK,
		Data:      model,
		Message:   "",
	}, nil
}

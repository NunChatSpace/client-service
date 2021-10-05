package entities

import (
	"gorm.io/gorm"
)

func (d Db) User() UserInterface {
	return userDb(d)
}

type userDb struct {
	gorm *gorm.DB
}

type UserInterface interface {
	Add(info UserModel) (UserModel, error)
	Get(u UserModel) (UserModel, error)
}

type UserModel struct {
	Model
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Type       string `json:"type"`
	ContactID  string `json:"contact_id"`
	CitizenID  string `json:"citizen_id"`
}

func (UserModel) TableName() string {
	return "users"
}

func (udb userDb) Add(info UserModel) (UserModel, error) {
	tx := udb.gorm.Create(&info)

	return info, tx.Error
}

func (udb userDb) Get(u UserModel) (UserModel, error) {
	user := UserModel{}
	udb.gorm.Model(&u).First(&user)

	return user, nil
}

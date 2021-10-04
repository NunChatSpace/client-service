package entities

import "gorm.io/gorm"

func (d Db) Farm() FarmInterface {
	return farmDb(d)
}

type farmDb struct {
	gorm *gorm.DB
}

type FarmInterface interface {
	Add(info FarmModel) (FarmModel, error)
	Get(uid string) (FarmModel, error)
}

type FarmModel struct {
	Model
	Name string `json:"name"`
}

func (FarmModel) TableName() string {
	return "farms"
}

func (fdb farmDb) Add(info FarmModel) (FarmModel, error) {
	tx := fdb.gorm.Create(&info)

	return info, tx.Error
}

func (fdb farmDb) Get(uid string) (FarmModel, error) {
	farm := FarmModel{}
	fdb.gorm.Find(&farm, "id IN ?", uid)

	return farm, nil
}

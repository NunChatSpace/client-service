package entities

import "gorm.io/gorm"

func (d Db) Cattle() CattleInterface {
	return cattleDb(d)
}

type cattleDb struct {
	gorm *gorm.DB
}

type CattleInterface interface {
	Add(info CattleModel) (CattleModel, error)
	AddMany(infos []CattleModel) (int, error)
	Get(id string) (CattleModel, error)
	GetByNames(names []string) ([]CattleModel, error)
}

type CattleModel struct {
	Model
	AnimalName    string `json:"animal_name"`
	DeviceID      string `json:"device_id"`
	BirthDate     string `json:"birth_date"`
	Sex           string `json:"sex"`
	Color         string `json:"color"`
	Breed         string `json:"breed"`
	BirthWeight   string `json:"birth_weight"`
	Chest         string `json:"chest"`
	WeaningWeight string `json:"weaning_weight"`
	BuildNumber   string `json:"build_number"`
	BreederID     string `json:"breeder_id"`
	OwnerID       string `json:"owner_id"`
}

func (CattleModel) TableName() string {
	return "roles"
}

func (cdb cattleDb) Add(info CattleModel) (CattleModel, error) {
	tx := cdb.gorm.Create(&info)

	return info, tx.Error
}

func (cdb cattleDb) AddMany(infos []CattleModel) (int, error) {
	tx := cdb.gorm.Create(&infos)

	return int(tx.RowsAffected), tx.Error
}

func (cdb cattleDb) Get(id string) (CattleModel, error) {
	cattle := CattleModel{}
	cdb.gorm.Find(&cattle, "id = ?", id)

	return cattle, nil
}

func (cdb cattleDb) GetByNames(names []string) ([]CattleModel, error) {
	cattles := []CattleModel{}
	cdb.gorm.Find(&cattles, "animal_name IN ?", names)

	return cattles, nil
}

package deliveries_cattles

type CattlesRegisterModel struct {
	OwnerCitizenID string        `json:"owner_citizen_id"`
	Cattles        []CattleModel `json:"cattles"`
}

type CattleModel struct {
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
	BreederName   string `json:"breeder_name"`
}

package deliveries_cattles

import (
	"fmt"
	"net/http"

	"github.com/NunChatSpace/client-service/internal/deliveries"
	"github.com/NunChatSpace/client-service/internal/entities"
	"github.com/NunChatSpace/client-service/internal/response_wrapper"
)

func AddCattles(db entities.DB, model CattlesRegisterModel) (response_wrapper.Model, error) {
	owner, err := db.User().Get(entities.UserModel{
		CitizenID: model.OwnerCitizenID,
	})
	if err != nil {
		return deliveries.InternalError(CattlesRegisterModel{}, err)
	}

	cattlesModel := cattlesPrep(model.Cattles, owner.ID, getBreederIDs(db, model.Cattles))
	if len(cattlesModel) == 0 {
		return response_wrapper.Model{
			ErrorCode: http.StatusOK,
			Message:   "0 rows affect",
		}, nil
	}

	rows, err := db.Cattle().AddMany(cattlesModel)
	if err != nil {
		return deliveries.InternalError(CattlesRegisterModel{}, err)
	}

	return response_wrapper.Model{
		ErrorCode: http.StatusOK,
		Data:      model,
		Message:   fmt.Sprintf("%d rows affect", rows),
	}, nil
}

func getBreederIDs(db entities.DB, models []CattleModel) map[string]string {
	result := make(map[string]string)
	var breederNames []string

	for _, model := range models {
		_, ok := result[model.BreederName]
		if !ok {
			breederNames = append(breederNames, model.BreederName)
			result[model.BreederName] = ""
		}
	}

	cattleModels, err := db.Cattle().GetByNames(breederNames)
	if err != nil {
		return result
	}

	for _, cattle := range cattleModels {
		result[cattle.AnimalName] = cattle.ID
	}

	return result
}

func cattlesPrep(models []CattleModel, ownerID string, breeders map[string]string) []entities.CattleModel {
	var result []entities.CattleModel

	for _, model := range models {
		result = append(result, entities.CattleModel{
			AnimalName:    model.AnimalName,
			DeviceID:      model.DeviceID,
			BirthDate:     model.BirthDate,
			Sex:           model.Sex,
			Color:         model.Color,
			Breed:         model.Breed,
			BirthWeight:   model.BirthWeight,
			Chest:         model.Chest,
			WeaningWeight: model.WeaningWeight,
			BuildNumber:   model.BuildNumber,
			BreederID:     breeders[model.BreederName],
			OwnerID:       ownerID,
		})
	}

	return result
}

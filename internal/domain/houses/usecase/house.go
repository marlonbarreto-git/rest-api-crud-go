package usecase

import (
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/houses/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

func (uc *useCase) GetHouse(id string) (*entities.House, *utils.Error) {
	return uc.repository.GetHouse(id)
}

func (uc *useCase) GetHouses(params *entities.HousesPageParams) (*entities.HousesPage, *utils.Error) {
	return uc.repository.GetHouses(params.Size, params.Page)
}

func (uc *useCase) CreateHouse(house *entities.HousePayload) (*entities.House, *utils.Error) {
	return uc.repository.CreateHouse(*house)
}

func (uc *useCase) UpdateHouse(house *entities.HousePayload) (*entities.House, *utils.Error) {
	return uc.repository.UpdateHouse(*house)
}

func (uc *useCase) DeleteHouse(id string) *utils.Error {
	return uc.repository.DeleteHouse(id)
}

package usecase

import (
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/houses/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/houses/repository"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/dependencies"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

type (
	House interface {
		GetHouse(id string) (*entities.House, *utils.Error)
		GetHouses(params *entities.HousesPageParams) (*entities.HousesPage, *utils.Error)
		CreateHouse(house *entities.HousePayload) (*entities.House, *utils.Error)
		UpdateHouse(house *entities.HousePayload) (*entities.House, *utils.Error)
		DeleteHouse(id string) *utils.Error
	}

	useCase struct {
		repository repository.House
	}
)

func NewUseCase(container *dependencies.Container) House {
	return &useCase{
		repository: repository.NewRepository(container),
	}
}

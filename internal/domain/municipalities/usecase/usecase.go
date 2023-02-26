package usecase

import (
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/municipalities/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/municipalities/repository"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/dependencies"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

type (
	Municipality interface {
		GetMunicipality(id int) (*entities.Municipality, *utils.Error)
		GetMunicipalities(params *entities.MunicipalitiesPageParams) (*entities.MunicipalitiesPage, *utils.Error)
		CreateMunicipality(municipality *entities.MunicipalityPayload) (*entities.Municipality, *utils.Error)
		UpdateMunicipality(municipality *entities.MunicipalityPayload) (*entities.Municipality, *utils.Error)
		DeleteMunicipality(id int) *utils.Error
	}

	useCase struct {
		repository repository.Municipality
	}
)

func NewUseCase(container *dependencies.Container) Municipality {
	return &useCase{
		repository: repository.NewRepository(container),
	}
}

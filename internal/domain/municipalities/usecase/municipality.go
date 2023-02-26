package usecase

import (
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/municipalities/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

func (uc *useCase) GetMunicipality(id int) (*entities.Municipality, *utils.Error) {
	return uc.repository.GetMunicipality(id)
}

func (uc *useCase) GetMunicipalities(params *entities.MunicipalitiesPageParams) (*entities.MunicipalitiesPage, *utils.Error) {
	return uc.repository.GetMunicipalities(params.Size, params.Page)
}

func (uc *useCase) CreateMunicipality(municipality *entities.MunicipalityPayload) (*entities.Municipality, *utils.Error) {
	return uc.repository.CreateMunicipality(*municipality)
}

func (uc *useCase) UpdateMunicipality(municipality *entities.MunicipalityPayload) (*entities.Municipality, *utils.Error) {
	return uc.repository.UpdateMunicipality(*municipality)
}

func (uc *useCase) DeleteMunicipality(id int) *utils.Error {
	return uc.repository.DeleteMunicipality(id)
}

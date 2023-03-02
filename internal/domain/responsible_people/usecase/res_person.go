package usecase

import (
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/responsible_people/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

func (uc *useCase) GetResponsible(id int) (*entities.Responsible, *utils.Error) {
	return uc.repository.GetResponsible(id)
}

func (uc *useCase) GetResponsibles(params *entities.ResponsiblesPageParams) (*entities.ResponsiblesPage, *utils.Error) {
	return uc.repository.GetResponsibles(params.Size, params.Page)
}

func (uc *useCase) CreateResponsible(responsible *entities.ResponsiblePayload) (*entities.Responsible, *utils.Error) {
	return uc.repository.CreateResponsible(*responsible)
}

func (uc *useCase) DeleteResponsible(id int) *utils.Error {
	return uc.repository.DeleteResponsible(id)
}

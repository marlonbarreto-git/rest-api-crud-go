package usecase

import (
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/people/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

func (uc *useCase) GetPerson(id int) (*entities.Person, *utils.Error) {
	return uc.repository.GetPerson(id)
}

func (uc *useCase) GetPeople(params *entities.PeoplePageParams) (*entities.PeoplePage, *utils.Error) {
	return uc.repository.GetPeople(params.Size, params.Page)
}

func (uc *useCase) CreatePerson(person *entities.PersonPayload) (*entities.Person, *utils.Error) {
	return uc.repository.CreatePerson(*person)
}

func (uc *useCase) UpdatePerson(person *entities.PersonPayload) (*entities.Person, *utils.Error) {
	return uc.repository.UpdatePerson(*person)
}

func (uc *useCase) DeletePerson(id int) *utils.Error {
	return uc.repository.DeletePerson(id)
}

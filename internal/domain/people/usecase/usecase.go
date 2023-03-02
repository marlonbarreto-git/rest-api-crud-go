package usecase

import (
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/people/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/people/repository"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/dependencies"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

type (
	Person interface {
		GetPerson(id int) (*entities.Person, *utils.Error)
		GetPeople(params *entities.PeoplePageParams) (*entities.PeoplePage, *utils.Error)
		CreatePerson(house *entities.PersonPayload) (*entities.Person, *utils.Error)
		UpdatePerson(house *entities.PersonPayload) (*entities.Person, *utils.Error)
		DeletePerson(id int) *utils.Error
	}

	useCase struct {
		repository repository.Person
	}
)

func NewUseCase(container *dependencies.Container) Person {
	return &useCase{
		repository: repository.NewRepository(container),
	}
}

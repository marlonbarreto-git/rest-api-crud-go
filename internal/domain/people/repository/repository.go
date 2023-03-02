package repository

import (
	"database/sql"

	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/people/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/dependencies"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

type (
	Person interface {
		GetPerson(id int) (*entities.Person, *utils.Error)
		GetPeople(size, page int) (*entities.PeoplePage, *utils.Error)
		CreatePerson(person entities.PersonPayload) (*entities.Person, *utils.Error)
		UpdatePerson(person entities.PersonPayload) (*entities.Person, *utils.Error)
		DeletePerson(id int) *utils.Error
	}

	repository struct {
		database *sql.DB
	}
)

func NewRepository(container *dependencies.Container) Person {
	return &repository{
		database: container.Database,
	}
}

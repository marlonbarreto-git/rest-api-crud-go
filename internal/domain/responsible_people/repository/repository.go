package repository

import (
	"database/sql"

	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/responsible_people/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/dependencies"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

type (
	Responsible interface {
		GetResponsible(id int) (*entities.Responsible, *utils.Error)
		GetResponsibles(size, page int) (*entities.ResponsiblesPage, *utils.Error)
		CreateResponsible(responsible entities.ResponsiblePayload) (*entities.Responsible, *utils.Error)
		DeleteResponsible(responsibleID, personID int) *utils.Error
	}

	repository struct {
		database *sql.DB
	}
)

func NewRepository(container *dependencies.Container) Responsible {
	return &repository{
		database: container.Database,
	}
}

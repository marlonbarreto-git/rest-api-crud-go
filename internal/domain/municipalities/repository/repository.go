package repository

import (
	"database/sql"

	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/municipalities/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/dependencies"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

type (
	Municipality interface {
		GetMunicipality(id int) (*entities.Municipality, *utils.Error)
		GetMunicipalities(size, page int) (*entities.MunicipalitiesPage, *utils.Error)
		CreateMunicipality(municipality entities.MunicipalityPayload) (*entities.Municipality, *utils.Error)
		UpdateMunicipality(municipality entities.MunicipalityPayload) (*entities.Municipality, *utils.Error)
		DeleteMunicipality(id int) *utils.Error
	}

	repository struct {
		database *sql.DB
	}
)

func NewRepository(container *dependencies.Container) Municipality {
	return &repository{
		database: container.Database,
	}
}

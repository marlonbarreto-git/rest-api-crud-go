package repository

import (
	"database/sql"

	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/houses/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/dependencies"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

type (
	House interface {
		GetHouse(id int) (*entities.House, *utils.Error)
		GetHouses(size, page int) (*entities.HousesPage, *utils.Error)
		CreateHouse(house entities.HousePayload) (*entities.House, *utils.Error)
		UpdateHouse(house entities.HousePayload) (*entities.House, *utils.Error)
		DeleteHouse(id int) *utils.Error
	}

	repository struct {
		database *sql.DB
	}
)

func NewRepository(container *dependencies.Container) House {
	return &repository{
		database: container.Database,
	}
}

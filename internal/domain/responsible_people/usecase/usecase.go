package usecase

import (
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/responsible_people/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/responsible_people/repository"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/dependencies"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

type (
	Responsible interface {
		GetResponsible(id int) (*entities.Responsible, *utils.Error)
		GetResponsibles(params *entities.ResponsiblesPageParams) (*entities.ResponsiblesPage, *utils.Error)
		CreateResponsible(responsible *entities.ResponsiblePayload) (*entities.Responsible, *utils.Error)
		DeleteResponsible(id int) *utils.Error
	}

	useCase struct {
		repository repository.Responsible
	}
)

func NewUseCase(container *dependencies.Container) Responsible {
	return &useCase{
		repository: repository.NewRepository(container),
	}
}

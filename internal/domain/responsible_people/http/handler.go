package http

import (
	"github.com/gin-gonic/gin"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/responsible_people/usecase"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/dependencies"
)

type (
	Handler interface {
		GetResponsibles(ctx *gin.Context)
		GetResponsibleById(ctx *gin.Context)
		CreateResponsible(ctx *gin.Context)
		DeleteResponsible(ctx *gin.Context)
	}

	handler struct {
		useCase usecase.Responsible
	}
)

const (
	ResponsibleIDParam = "responsibleId"
	PersonIDParam      = "personId"
)

func NewHandler(container *dependencies.Container) Handler {
	return &handler{
		useCase: usecase.NewUseCase(container),
	}
}

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/people/usecase"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/dependencies"
)

type (
	Handler interface {
		GetPeople(ctx *gin.Context)
		GetPersonById(ctx *gin.Context)
		CreatePerson(ctx *gin.Context)
		UpdatePerson(ctx *gin.Context)
		DeletePerson(ctx *gin.Context)
	}

	handler struct {
		useCase usecase.Person
	}
)

const IdParam = "id"

func NewHandler(container *dependencies.Container) Handler {
	return &handler{
		useCase: usecase.NewUseCase(container),
	}
}

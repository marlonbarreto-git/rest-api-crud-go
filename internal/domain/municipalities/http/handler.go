package http

import (
	"github.com/gin-gonic/gin"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/municipalities/usecase"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/dependencies"
)

type (
	Handler interface {
		GetMunicipalities(ctx *gin.Context)
		GetMunicipalityById(ctx *gin.Context)
		CreateMunicipality(ctx *gin.Context)
		UpdateMunicipality(ctx *gin.Context)
		DeleteMunicipality(ctx *gin.Context)
	}

	handler struct {
		useCase usecase.Municipality
	}
)

const IdParam = "id"

func NewHandler(container *dependencies.Container) Handler {
	return &handler{
		useCase: usecase.NewUseCase(container),
	}
}

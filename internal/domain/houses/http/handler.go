package http

import (
	"github.com/gin-gonic/gin"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/houses/usecase"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/dependencies"
)

type (
	Handler interface {
		GetHouses(ctx *gin.Context)
		GetHouseById(ctx *gin.Context)
		CreateHouse(ctx *gin.Context)
		UpdateHouse(ctx *gin.Context)
		DeleteHouse(ctx *gin.Context)
	}

	handler struct {
		useCase usecase.House
	}
)

const IdParam = "id"

func NewHandler(container *dependencies.Container) Handler {
	return &handler{
		useCase: usecase.NewUseCase(container),
	}
}

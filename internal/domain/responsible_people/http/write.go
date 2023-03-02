package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/responsible_people/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

func (handler *handler) CreateResponsible(ctx *gin.Context) {
	responsible := &entities.ResponsiblePayload{}
	if err := ctx.ShouldBindJSON(&responsible); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	createdResponsible, err := handler.useCase.CreateResponsible(responsible)
	if err != nil {
		statusCode := utils.ConvertMessageToCode(string(err.Message))
		ctx.JSON(statusCode, err)
		return
	}

	ctx.JSON(http.StatusCreated, createdResponsible)
}

func (handler *handler) DeleteResponsible(ctx *gin.Context) {
	pathParams := &entities.ResponsibleParams{}
	if err := ctx.ShouldBindUri(&pathParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := handler.useCase.DeleteResponsible(pathParams.Id); err != nil {
		statusCode := utils.ConvertMessageToCode(string(err.Message))
		ctx.JSON(statusCode, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

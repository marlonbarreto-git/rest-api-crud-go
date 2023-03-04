package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/responsible_people/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

func (handler *handler) GetResponsibles(ctx *gin.Context) {
	queryParams := &entities.ResponsiblesPageParams{}
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	responsibles, err := handler.useCase.GetResponsibles(queryParams)
	if err != nil {
		statusCode := utils.ConvertMessageToCode(string(err.Message))
		ctx.JSON(statusCode, err)
		return
	}

	ctx.JSON(http.StatusOK, responsibles)
}

func (handler *handler) GetResponsibleById(ctx *gin.Context) {
	pathParams := &entities.ResponsibleParams{}
	if err := ctx.ShouldBindUri(&pathParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	responsible, err := handler.useCase.GetResponsible(pathParams.ResponsibleID)
	if err != nil {
		statusCode := utils.ConvertMessageToCode(string(err.Message))
		ctx.JSON(statusCode, err)
		return
	}

	ctx.JSON(http.StatusOK, responsible)
}

package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/municipalities/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

func (handler *handler) GetMunicipalities(ctx *gin.Context) {
	queryParams := &entities.MunicipalitiesPageParams{}
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	municipalities, err := handler.useCase.GetMunicipalities(queryParams)
	if err != nil {
		statusCode := utils.ConvertMessageToCode(string(err.Message))
		ctx.JSON(statusCode, err)
		return
	}

	ctx.JSON(http.StatusOK, municipalities)
}

func (handler *handler) GetMunicipalityById(ctx *gin.Context) {
	pathParams := &entities.MunicipalityParams{}
	if err := ctx.ShouldBindUri(&pathParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	municipality, err := handler.useCase.GetMunicipality(pathParams.Id)
	if err != nil {
		statusCode := utils.ConvertMessageToCode(string(err.Message))
		ctx.JSON(statusCode, err)
		return
	}

	ctx.JSON(http.StatusOK, municipality)
}

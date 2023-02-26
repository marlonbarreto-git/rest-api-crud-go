package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/municipalities/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

func (handler *handler) CreateMunicipality(ctx *gin.Context) {
	municipality := &entities.MunicipalityPayload{}
	if err := ctx.ShouldBindJSON(&municipality); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	createdMunicipality, err := handler.useCase.CreateMunicipality(municipality)
	if err != nil {
		statusCode := utils.ConvertMessageToCode(string(err.Message))
		ctx.JSON(statusCode, err)
		return
	}

	ctx.JSON(http.StatusCreated, createdMunicipality)
}

func (handler *handler) UpdateMunicipality(ctx *gin.Context) {
	pathParams := &entities.MunicipalityParams{}
	if err := ctx.ShouldBindUri(&pathParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	municipality := &entities.MunicipalityPayload{}
	if err := ctx.BindJSON(&municipality); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	municipality.Id = &pathParams.Id

	updatedMunicipality, err := handler.useCase.UpdateMunicipality(municipality)
	if err != nil {
		statusCode := utils.ConvertMessageToCode(string(err.Message))
		ctx.JSON(statusCode, err)
		return
	}

	ctx.JSON(http.StatusOK, updatedMunicipality)
}

func (handler *handler) DeleteMunicipality(ctx *gin.Context) {
	pathParams := &entities.MunicipalityParams{}
	if err := ctx.ShouldBindUri(&pathParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := handler.useCase.DeleteMunicipality(pathParams.Id); err != nil {
		statusCode := utils.ConvertMessageToCode(string(err.Message))
		ctx.JSON(statusCode, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

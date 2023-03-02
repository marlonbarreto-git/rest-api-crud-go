package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/houses/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

func (handler *handler) CreateHouse(ctx *gin.Context) {
	house := &entities.HousePayload{}
	if err := ctx.ShouldBindJSON(&house); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	createdHouse, err := handler.useCase.CreateHouse(house)
	if err != nil {
		statusCode := utils.ConvertMessageToCode(string(err.Message))
		ctx.JSON(statusCode, err)
		return
	}

	ctx.JSON(http.StatusCreated, createdHouse)
}

func (handler *handler) UpdateHouse(ctx *gin.Context) {
	pathParams := &entities.HouseParams{}
	if err := ctx.ShouldBindUri(&pathParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	house := &entities.HousePayload{}
	if err := ctx.BindJSON(&house); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	house.Id = &pathParams.Id

	updatedHouse, err := handler.useCase.UpdateHouse(house)
	if err != nil {
		statusCode := utils.ConvertMessageToCode(string(err.Message))
		ctx.JSON(statusCode, err)
		return
	}

	ctx.JSON(http.StatusOK, updatedHouse)
}

func (handler *handler) DeleteHouse(ctx *gin.Context) {
	pathParams := &entities.HouseParams{}
	if err := ctx.ShouldBindUri(&pathParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := handler.useCase.DeleteHouse(pathParams.Id); err != nil {
		statusCode := utils.ConvertMessageToCode(string(err.Message))
		ctx.JSON(statusCode, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

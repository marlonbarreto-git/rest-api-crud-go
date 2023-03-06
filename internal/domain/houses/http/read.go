package http

import (
	"github.com/gin-gonic/gin"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/houses/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
	"net/http"
)

func (handler *handler) GetHouses(ctx *gin.Context) {
	queryParams := &entities.HousesPageParams{}
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	houses, err := handler.useCase.GetHouses(queryParams)
	if err != nil {
		statusCode := utils.ConvertMessageToCode(string(err.Message))
		ctx.JSON(statusCode, err)
		return
	}

	ctx.JSON(http.StatusOK, houses)
}

func (handler *handler) GetHouseById(ctx *gin.Context) {
	pathParams := &entities.HouseParams{}
	if err := ctx.ShouldBindUri(&pathParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	house, err := handler.useCase.GetHouse(pathParams.Id)
	if err != nil {
		statusCode := utils.ConvertMessageToCode(string(err.Message))
		ctx.JSON(statusCode, err)
		return
	}

	ctx.JSON(http.StatusOK, house)
}

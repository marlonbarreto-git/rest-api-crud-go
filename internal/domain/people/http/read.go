package http

import (
	"github.com/gin-gonic/gin"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/people/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
	"net/http"
)

func (handler *handler) GetPeople(ctx *gin.Context) {
	queryParams := &entities.PeoplePageParams{}
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	people, err := handler.useCase.GetPeople(queryParams)
	if err != nil {
		statusCode := utils.ConvertMessageToCode(string(err.Message))
		ctx.JSON(statusCode, err)
		return
	}

	ctx.JSON(http.StatusOK, people)
}

func (handler *handler) GetPersonById(ctx *gin.Context) {
	pathParams := &entities.PersonParams{}
	if err := ctx.ShouldBindUri(&pathParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	person, err := handler.useCase.GetPerson(pathParams.Id)
	if err != nil {
		statusCode := utils.ConvertMessageToCode(string(err.Message))
		ctx.JSON(statusCode, err)
		return
	}

	ctx.JSON(http.StatusOK, person)
}

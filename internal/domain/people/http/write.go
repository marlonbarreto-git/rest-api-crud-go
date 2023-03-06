package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/people/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

func (handler *handler) CreatePerson(ctx *gin.Context) {
	person := &entities.PersonPayload{}
	if err := ctx.ShouldBindJSON(&person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	createdPerson, err := handler.useCase.CreatePerson(person)
	if err != nil {
		statusCode := utils.ConvertMessageToCode(string(err.Message))
		ctx.JSON(statusCode, err)
		return
	}

	ctx.JSON(http.StatusCreated, createdPerson)
}

func (handler *handler) UpdatePerson(ctx *gin.Context) {
	pathParams := &entities.PersonParams{}
	if err := ctx.ShouldBindUri(&pathParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	person := &entities.PersonPayload{}
	if err := ctx.BindJSON(&person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	person.Id = &pathParams.Id

	updatedPerson, err := handler.useCase.UpdatePerson(person)
	if err != nil {
		statusCode := utils.ConvertMessageToCode(string(err.Message))
		ctx.JSON(statusCode, err)
		return
	}

	ctx.JSON(http.StatusOK, updatedPerson)
}

func (handler *handler) DeletePerson(ctx *gin.Context) {
	pathParams := &entities.PersonParams{}
	if err := ctx.ShouldBindUri(&pathParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := handler.useCase.DeletePerson(pathParams.Id); err != nil {
		statusCode := utils.ConvertMessageToCode(string(err.Message))
		ctx.JSON(statusCode, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

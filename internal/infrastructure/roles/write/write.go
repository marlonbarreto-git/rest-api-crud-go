package write

import (
	"fmt"

	"github.com/gin-gonic/gin"

	housesHttp "github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/houses/http"
	municipalitiesHttp "github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/municipalities/http"
	peopleHttp "github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/people/http"
	responsibleHttp "github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/responsible_people/http"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/dependencies"
)

type Write struct {
	container *dependencies.Container
}

func NewWrite(container *dependencies.Container) *Write {
	return &Write{container}
}

func (write *Write) RegisterRoutes(basePath string) func(group *gin.RouterGroup) {
	municipalityHandler := municipalitiesHttp.NewHandler(write.container)
	houseHandler := housesHttp.NewHandler(write.container)
	personHandler := peopleHttp.NewHandler(write.container)
	responsibleHandler := responsibleHttp.NewHandler(write.container)

	return func(g *gin.RouterGroup) {
		v1Group := g.Group(basePath + "/v1")
		roleGroup := v1Group.Group("/write")

		roleGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		municipalitiesGroup := roleGroup.Group("/municipalities")
		municipalitiesGroup.POST("", municipalityHandler.CreateMunicipality)
		municipalitiesGroup.PATCH(fmt.Sprintf("/:%s", municipalitiesHttp.IdParam), municipalityHandler.UpdateMunicipality)
		municipalitiesGroup.DELETE(fmt.Sprintf("/:%s", municipalitiesHttp.IdParam), municipalityHandler.DeleteMunicipality)

		housesGroup := roleGroup.Group("/houses")
		housesGroup.POST("", houseHandler.CreateHouse)
		housesGroup.PATCH(fmt.Sprintf("/:%s", housesHttp.IdParam), houseHandler.UpdateHouse)
		housesGroup.DELETE(fmt.Sprintf("/:%s", housesHttp.IdParam), houseHandler.DeleteHouse)

		peopleGroup := roleGroup.Group("/people")
		peopleGroup.POST("", personHandler.CreatePerson)
		peopleGroup.PATCH(fmt.Sprintf("/:%s", peopleHttp.IdParam), personHandler.UpdatePerson)
		peopleGroup.DELETE(fmt.Sprintf("/:%s", peopleHttp.IdParam), personHandler.DeletePerson)

		responsiblesGroup := roleGroup.Group("/responsibles")
		responsiblesGroup.POST("", responsibleHandler.CreateResponsible)
		responsiblesGroup.DELETE(fmt.Sprintf("/:%s/:%s", responsibleHttp.ResponsibleIDParam, responsibleHttp.PersonIDParam), responsibleHandler.DeleteResponsible)
	}
}

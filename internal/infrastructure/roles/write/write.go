package write

import (
	"fmt"
	phttp "github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/people/http"
	rhttp "github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/responsible_people/http"

	"github.com/gin-gonic/gin"

	hhttp "github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/houses/http"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/municipalities/http"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/dependencies"
)

type Write struct {
	container *dependencies.Container
}

func NewWrite(container *dependencies.Container) *Write {
	return &Write{container}
}

func (write *Write) RegisterRoutes(basePath string) func(group *gin.RouterGroup) {
	municipalityHandler := http.NewHandler(write.container)
	houseHandler := hhttp.NewHandler(write.container)
	personHandler := phttp.NewHandler(write.container)
	responsibleHandler := rhttp.NewHandler(write.container)

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
		municipalitiesGroup.PATCH(fmt.Sprintf("/:%s", http.IdParam), municipalityHandler.UpdateMunicipality)
		municipalitiesGroup.DELETE(fmt.Sprintf("/:%s", http.IdParam), municipalityHandler.DeleteMunicipality)

		housesGroup := roleGroup.Group("/houses")
		housesGroup.POST("", houseHandler.CreateHouse)
		housesGroup.PATCH(fmt.Sprintf("/:%s", hhttp.IdParam), houseHandler.UpdateHouse)
		housesGroup.DELETE(fmt.Sprintf("/:%s", hhttp.IdParam), houseHandler.DeleteHouse)

		peopleGroup := roleGroup.Group("/people")
		peopleGroup.POST("", personHandler.CreatePerson)
		peopleGroup.PATCH(fmt.Sprintf("/:%s", phttp.IdParam), personHandler.UpdatePerson)
		peopleGroup.DELETE(fmt.Sprintf("/:%s", phttp.IdParam), personHandler.DeletePerson)

		responsiblesGroup := roleGroup.Group("/responsibles")
		responsiblesGroup.POST("", responsibleHandler.CreateResponsible)
		responsiblesGroup.DELETE(fmt.Sprintf("/:%s", rhttp.IdParam), responsibleHandler.DeleteResponsible)
	}
}

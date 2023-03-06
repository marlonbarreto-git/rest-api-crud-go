package read

import (
	"github.com/gin-gonic/gin"
	hhttp "github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/houses/http"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/municipalities/http"
	phttp "github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/people/http"
	rhttp "github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/responsible_people/http"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/dependencies"
)

type read struct {
	container *dependencies.Container
}

func NewRead(container *dependencies.Container) *read {
	return &read{container}
}

func (read *read) RegisterRoutes(basePath string) func(group *gin.RouterGroup) {
	municipalityHandler := http.NewHandler(read.container)
	houseHandler := hhttp.NewHandler(read.container)
	personHandler := phttp.NewHandler(read.container)
	responsibleHandler := rhttp.NewHandler(read.container)

	return func(g *gin.RouterGroup) {
		v1Group := g.Group(basePath + "/v1")
		roleGroup := v1Group.Group("/read")

		roleGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		municipalitiesGroup := roleGroup.Group("/municipalities")
		municipalitiesGroup.GET("/", municipalityHandler.GetMunicipalities)
		municipalitiesGroup.GET("/:id", municipalityHandler.GetMunicipalityById)

		housesGroup := roleGroup.Group("/houses")
		housesGroup.GET("/", houseHandler.GetHouses)
		housesGroup.GET("/:id", houseHandler.GetHouseById)

		peopleGroup := roleGroup.Group("/people")
		peopleGroup.GET("/", personHandler.GetPeople)
		peopleGroup.GET("/:id", personHandler.GetPersonById)

		responsiblesGroup := roleGroup.Group("/responsibles")
		responsiblesGroup.GET("/", responsibleHandler.GetResponsibles)
		responsiblesGroup.GET("/:id", responsibleHandler.GetResponsibleById)
	}
}

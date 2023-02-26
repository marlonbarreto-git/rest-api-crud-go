package read

import (
	"github.com/gin-gonic/gin"

	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/municipalities/http"
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
	}
}

package write

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/municipalities/http"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/dependencies"
)

type write struct {
	container *dependencies.Container
}

func NewWrite(container *dependencies.Container) *write {
	return &write{container}
}

func (write *write) RegisterRoutes(basePath string) func(group *gin.RouterGroup) {
	municipalityHandler := http.NewHandler(write.container)

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
	}
}

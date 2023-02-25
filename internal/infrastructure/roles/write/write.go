package write

import (
	"github.com/gin-gonic/gin"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/dependencies"
)

type write struct {
	container *dependencies.Container
}

func NewWrite(container *dependencies.Container) *write {
	return &write{container}
}

func (write *write) RegisterRoutes(basePath string) func(group *gin.RouterGroup) {
	return func(g *gin.RouterGroup) {
		v1Group := g.Group(basePath + "/v1")
		roleGroup := v1Group.Group("/write")

		roleGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}

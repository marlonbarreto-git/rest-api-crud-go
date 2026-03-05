package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/context"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/roles"
)

var buildContext = infrastructure.NewContextBuilder()

func main() {
	ctx := context.InitializeContext()

	server := &infrastructure.Server{}
	gin.SetMode(gin.DebugMode)
	server.Engine = gin.Default()

	server.Engine.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Referrer-Policy", "no-referrer-when-downgrade")
		c.Header("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'; sandbox")
		c.Header("Content-Security-Policy", "connect-src 'self' http://localhost:8080")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}
		c.Next()
	})

	server.Engine.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Content-Type", "Authorization"},
	}))

	port := ctx.Port()

	group := server.Group("")

	configure := getRoleConfiguratorFunction(buildContext.GetRoutingGroup(), group, ctx.Role().IsAll())

	for _, role := range getRoles(ctx) {
		err := configure(role)
		if err != nil {
			panic(interface{}(err))
		}
	}

	if err := server.Run(":" + port); err != nil {
		panic(interface{}(err))
	}
}

func getRoles(ctx context.AppContext) []roles.Role {
	if ctx.Role().IsAll() {
		return roles.GetAllRoles()
	}

	return []roles.Role{ctx.Role()}
}

func getRoleConfiguratorFunction(routes infrastructure.RoutingGroup, group *gin.RouterGroup, isAll bool) func(role roles.Role) error {
	fmt.Printf("routes %v group %v, isAll %b", routes, group, isAll)
	return func(role roles.Role) error {
		if routesRole, ok := routes[role]; ok {
			routesRole()(group)

			return nil
		}

		if isAll {
			return nil
		}

		return fmt.Errorf("given routes does not contain endpoints for the \"%s\" application role", role)
	}
}

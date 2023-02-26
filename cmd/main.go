package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

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

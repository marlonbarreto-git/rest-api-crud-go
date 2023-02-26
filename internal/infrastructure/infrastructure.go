package infrastructure

import (
	"github.com/gin-gonic/gin"

	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/dependencies"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/roles"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/roles/read"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/roles/write"
)

type (
	BuildContext interface {
		GetRoutingGroup() RoutingGroup
	}

	RoutingGroup map[roles.Role]func() func(*gin.RouterGroup)

	Server struct {
		*gin.Engine
	}

	contextBuilder struct{}
)

const basePath = "/"

func NewContextBuilder() BuildContext {
	return &contextBuilder{}
}

func (contextBuilder) GetRoutingGroup() RoutingGroup {
	container := dependencies.StartDependencies()

	return RoutingGroup{
		roles.RoleRead: func() func(*gin.RouterGroup) {
			return read.NewRead(container).RegisterRoutes(basePath)
		},
		roles.RoleWrite: func() func(*gin.RouterGroup) {
			return write.NewWrite(container).RegisterRoutes(basePath)
		},
	}
}

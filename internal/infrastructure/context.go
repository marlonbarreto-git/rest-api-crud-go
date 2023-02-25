package infrastructure

import (
	"os"

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

	AppContext interface {
		ApplicationName() string
		Environment() Environment
		Role() roles.Role
		Port() string
	}

	RoutingGroup map[roles.Role]func() func(*gin.RouterGroup)

	Server struct {
		*gin.Engine
	}

	contextBuilder struct{}
)

const basePath = "/"

var (
	contextInitialized bool
	applicationName    string
	environment        Environment
	role               roles.Role
	port               = "8080"
)

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

func InitializeContext() AppContext {
	if contextInitialized {
		panic(interface{}("application context not initialized"))
	}

	applicationName = os.Getenv("APPLICATION")

	if envPort := os.Getenv("PORT"); len(envPort) != 0 {
		port = envPort
	}

	env, err := ConvertEnvironment(os.Getenv("ENVIRONMENT"))
	if err != nil {
		panic(interface{}(err))
	}

	environment = *env

	rawRole, err := roles.ConvertRole(os.Getenv("ROLE"))
	if err != nil {
		panic(interface{}(err))
	}

	role = *rawRole

	contextInitialized = true

	return &contextBuilder{}
}

func (ctx *contextBuilder) ApplicationName() string {
	return applicationName
}

func (ctx *contextBuilder) Environment() Environment {
	return environment
}

func (ctx *contextBuilder) Role() roles.Role {
	return role
}

func (ctx *contextBuilder) Port() string {
	return port
}

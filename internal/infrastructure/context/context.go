package context

import (
	"os"

	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/roles"
)

type (
	AppContext interface {
		ApplicationName() string
		Environment() roles.Environment
		Role() roles.Role
		Port() string
	}

	appContext struct{}
)

var (
	contextInitialized bool
	applicationName    string
	environment        roles.Environment
	role               roles.Role
	port               = "8080"
)

func InitializeContext() AppContext {
	if contextInitialized {
		panic(interface{}("application context not initialized"))
	}

	applicationName = os.Getenv("APPLICATION")

	if envPort := os.Getenv("PORT"); len(envPort) != 0 {
		port = envPort
	}

	env, err := roles.ConvertEnvironment(os.Getenv("ENVIRONMENT"))
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

	return &appContext{}
}

func GetContext() AppContext {
	if !contextInitialized {
		panic("application context not initialized")
	}

	return &appContext{}
}

func (ctx *appContext) ApplicationName() string {
	return applicationName
}

func (ctx *appContext) Environment() roles.Environment {
	return environment
}

func (ctx *appContext) Role() roles.Role {
	return role
}

func (ctx *appContext) Port() string {
	return port
}

package configurations

import (
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/configurations/environments/develop"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/configurations/environments/local"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/configurations/environments/production"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/context"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/roles"
)

type Configuration interface {
	GetDBConfig() *entities.DBConfig
}

func GetConfiguration() Configuration {
	switch context.GetContext().Environment() {
	case roles.EnvProduction:
		return production.NewConfiguration()
	case roles.EnvDevelop:
		return develop.NewConfiguration()
	default:
		return local.NewConfiguration()
	}
}

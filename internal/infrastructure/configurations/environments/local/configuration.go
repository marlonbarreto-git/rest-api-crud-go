package local

import "github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/entities"

type local struct{}

func NewConfiguration() *local {
	return &local{}
}

func (config *local) GetDBConfig() *entities.DBConfig {
	return &entities.DBConfig{
		ConnectionPath: "./local-crud.db",
	}
}

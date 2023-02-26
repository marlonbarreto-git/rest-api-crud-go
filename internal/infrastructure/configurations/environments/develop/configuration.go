package develop

import "github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/entities"

type develop struct{}

func NewConfiguration() *develop {
	return &develop{}
}

func (config *develop) GetDBConfig() *entities.DBConfig {
	return nil
}

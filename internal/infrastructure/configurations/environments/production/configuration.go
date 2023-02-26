package production

import "github.com/marlonbarreto-git/rest-api-crud-go/internal/infrastructure/entities"

type production struct{}

func NewConfiguration() *production {
	return &production{}
}

func (config *production) GetDBConfig() *entities.DBConfig {
	return nil
}

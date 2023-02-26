package entities

import "github.com/marlonbarreto-git/rest-api-crud-go/utils"

type (
	Municipality struct {
		Id   *int    `json:"id"`
		Name *string `json:"name"`
	}

	MunicipalitiesPage struct {
		utils.Page
		Data Municipalities `json:"data"`
	}

	Municipalities []Municipality
)

func (m *Municipality) IsEmpty() bool {
	return m.Id == nil || m.Name == nil
}

package entities

import "github.com/marlonbarreto-git/rest-api-crud-go/utils"

type (
	House struct {
		Id             *int    `json:"id"`
		Address        *string `json:"address"`
		IdOwner        *int    `json:"id_owner"`
		IdMunicipality *int    `json:"id_municipality"`
	}

	HousesPage struct {
		utils.Page
		Data Houses `json:"data"`
	}

	Houses []House
)

func (m *House) IsEmpty() bool {
	return m.Id == nil || m.Address == nil || m.IdOwner == nil || m.IdMunicipality == nil
}

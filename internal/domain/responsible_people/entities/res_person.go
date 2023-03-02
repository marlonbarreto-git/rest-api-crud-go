package entities

import "github.com/marlonbarreto-git/rest-api-crud-go/utils"

type (
	Responsible struct {
		IdResponsible *int `json:"idResponsible"`
		IdPerson      *int `json:"idPerson"`
	}

	ResponsiblesPage struct {
		utils.Page
		Data Responsibles `json:"data"`
	}

	Responsibles []Responsible
)

func (r *Responsible) IsEmpty() bool {
	return r.IdResponsible == nil || r.IdPerson == nil
}

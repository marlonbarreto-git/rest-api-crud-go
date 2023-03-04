package entities

import "github.com/marlonbarreto-git/rest-api-crud-go/utils"

type (
	Person struct {
		Id        *int    `json:"id"`
		Forename  *string `json:"forename"`
		Surname   *string `json:"surname"`
		BirthDate *string `json:"birthDate"`
		Sex       *string `json:"sex"`
		IdHome    *string `json:"idHome"`
	}

	PeoplePage struct {
		utils.Page
		Data People `json:"data"`
	}

	People []Person
)

func (p *Person) IsEmpty() bool {
	return p.Id == nil || p.Forename == nil || p.Surname == nil || p.BirthDate == nil || p.IdHome == nil
}

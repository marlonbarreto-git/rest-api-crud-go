package entities

type PersonPayload struct {
	Id        *int    `json:"id"`
	Forename  *string `json:"forename"`
	Surname   *string `json:"surname"`
	BirthDate *string `json:"birthDate"`
	Sex       *string `json:"sex"`
	IdHome    *int    `json:"idHome"`
}

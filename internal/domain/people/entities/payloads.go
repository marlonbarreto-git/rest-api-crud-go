package entities

type PersonPayload struct {
	Id        *int    `json:"id" binding:"required"`
	Forename  *string `json:"forename" binding:"required,alpha"`
	Surname   *string `json:"surname" binding:"required,alpha"`
	BirthDate *string `json:"birthDate" binding:"required"`
	Sex       *string `json:"sex" binding:"required,oneof=M F"`
	IdHome    *string `json:"idHome" binding:"omitempty"`
}

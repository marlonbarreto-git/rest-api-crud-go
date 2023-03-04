package entities

type ResponsiblePayload struct {
	IdResponsible *int `json:"idResponsible" binding:"required"`
	IdPerson      *int `json:"idPerson" binding:"required"`
}

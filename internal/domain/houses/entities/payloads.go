package entities

type HousePayload struct {
	Id             *string `json:"id" binding:"required"`
	Address        *string `json:"address" binding:"required"`
	IdOwner        *int    `json:"idOwner" binding:"required"`
	IdMunicipality *int    `json:"idMunicipality" binding:"required"`
}

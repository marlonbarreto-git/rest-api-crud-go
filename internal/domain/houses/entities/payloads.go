package entities

type HousePayload struct {
	Id             *int    `json:"id"`
	Address        *string `json:"address"`
	IdOwner        *int    `json:"idOwner"`
	IdMunicipality *int    `json:"idMunicipality"`
}

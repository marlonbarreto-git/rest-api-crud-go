package entities

type HousePayload struct {
	Id             *int    `json:"id"`
	Address        *string `json:"address"`
	IdOwner        *int    `json:"id_owner"`
	IdMunicipality *int    `json:"id_municipality"`
}

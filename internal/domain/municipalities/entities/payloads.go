package entities

type MunicipalityPayload struct {
	Id   *int    `json:"-"`
	Name *string `json:"name"`
}

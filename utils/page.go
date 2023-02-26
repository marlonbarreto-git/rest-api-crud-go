package utils

type Page struct {
	Size  int `json:"size"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Count int `json:"count"`
}

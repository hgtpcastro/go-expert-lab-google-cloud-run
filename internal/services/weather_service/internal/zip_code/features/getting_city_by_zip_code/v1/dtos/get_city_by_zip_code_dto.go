package dtos

type GetCityByZipCodeResponseDTO struct {
	Localidade string `json:"localidade,omitempty"`
	Erro       bool   `json:"erro,omitempty"`
}

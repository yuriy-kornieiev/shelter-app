package dto

import "yaTrivoga/models"

type CityDTO struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func NewCityDTO(city models.City) CityDTO {
	dto := CityDTO{
		Id:   city.Id.Int64,
		Name: city.Name.String,
	}
	return dto
}

package dto

type CreateUpdateContentDto struct {
	Name   string `json:"name" validate:"required"`
	UserId string `json:"userId" validate:"required"`
}

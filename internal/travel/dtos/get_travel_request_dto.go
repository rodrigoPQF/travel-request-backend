package dtos

type GetTravelRequestInputDto struct {
	Id string `validate:"required,uuid"`
}


package dtos

type GetTravelRequestInputDto struct {
	Id string `validate:"required,uuid" example:"f57e8039-c69d-4de6-b483-eab41c804d16"`
}


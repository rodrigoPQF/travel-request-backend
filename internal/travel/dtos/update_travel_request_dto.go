package dtos

type UpdateTravelRequestBodyInputDto struct {
	Status string `json:"status" validate:"required,oneof=APPROVED CANCELED"`
}

type UpdateTravelRequestParamsInputDto struct {
	Id string `validate:"required,uuid" example:"f57e8039-c69d-4de6-b483-eab41c804d16"`
}

type UpdateTravelRequestInputDto struct {
	UpdateTravelRequestBodyInputDto
	UpdateTravelRequestParamsInputDto
}
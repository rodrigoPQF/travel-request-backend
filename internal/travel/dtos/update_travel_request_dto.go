package dtos

type UpdateTravelRequestBodyInputDto struct {
	Status string `json:"status" validate:"required,oneof=APPROVED CANCELED"`
}

type UpdateTravelRequestParamsInputDto struct {
	Id string `validate:"required,uuid"`
}

type UpdateTravelRequestInputDto struct {
	UpdateTravelRequestBodyInputDto
	UpdateTravelRequestParamsInputDto
}
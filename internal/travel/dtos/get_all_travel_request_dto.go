package dtos

type GetAllTravelRequestInputDto struct {
	Status string `form:"status" validate:"omitempty,oneof=REQUESTED APPROVED CANCELED" copier:"Status"`
	Page  int `form:"page" validate:"omitempty,min=1" copier:"Page" example:"1"`
	Limit int `form:"limit" validate:"omitempty,min=1" copier:"Limit" example:"2"`  
}


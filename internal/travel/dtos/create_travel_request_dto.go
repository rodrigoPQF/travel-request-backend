package dtos



type CreateTravelRequestInputDto struct {
	Id string `json:"id" validate:"required,uuid" copier:"Id" example:"f57e8039-c69d-4de6-b483-eab41c804d16"`
	ApplicantName string `json:"applicantName" validate:"required,min=3,max=100" copier:"ApplicantName" example:"Raul"`
	Destination string `json:"destination" validate:"required,min=3,max=100" copier:"Destination" example:"Rio de Janeiro"`
	DepartureDate string `json:"departureDate" validate:"required,datetime=2006-01-02" copier:"DepartureDate" example:"2001-01-30"`
	ReturnDate string `json:"returnDate" validate:"required,datetime=2006-01-02" copier:"ReturnDate" example:"2001-06-30"`
	Status string `json:"status" validate:"required,oneof=REQUESTED APPROVED CANCELED"`
}

type CreateTravelRequestOutputDto struct {
}



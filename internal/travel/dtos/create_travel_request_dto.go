package dtos

type CreateTravelRequestInputDto struct {
	Id string `json:"id" validate:"required,uuid" copier:"Id"`
	ApplicantName string `json:"applicantName" validate:"required,min=3,max=100" copier:"ApplicantName"`
	Destination string `json:"destination" validate:"required,min=3,max=100" copier:"Destination"`
	DepartureDate string `json:"departureDate" validate:"required,datetime=2006-01-02" copier:"DepartureDate"`
	ReturnDate string `json:"returnDate" validate:"required,datetime=2006-01-02" copier:"ReturnDate"`
	Status string `json:"status" validate:"required,oneof=REQUESTED APPROVED CANCELED"`
}

type CreateTravelRequestOutputDto struct {
}



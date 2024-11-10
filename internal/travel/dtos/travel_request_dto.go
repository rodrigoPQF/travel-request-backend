package dtos

type TravelRequestDto struct {
	Id string `json:"id,omitempty" example:"f57e8039-c69d-4de6-b483-eab41c804d16"`
	ApplicantName string `json:"applicantName,omitempty" example:"Raul"`
	Destination string `json:"destination,omitempty" example:"Rio de Janeiro"`
	DepartureDate string `json:"departureDate,omitempty" example:"2001-01-30"`
	ReturnDate string `json:"returnDate,omitempty" example:"2002-06-30"`
	Status string `json:"status,omitempty" example:"REQUESTED"`
}
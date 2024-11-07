package dtos

type TravelRequestDto struct {
	Id string `json:"id,omitempty"`
	ApplicantName string `json:"applicantName,omitempty"`
	Destination string `json:"destination,omitempty"`
	DepartureDate string `json:"departureDate,omitempty"`
	ReturnDate string `json:"returnDate,omitempty"`
	Status string `json:"status,omitempty"`
}
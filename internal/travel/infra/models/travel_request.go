package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	Requested   Status = "REQUESTED"
	Approved  Status = "APPROVED"
	Canceled  Status = "CANCELED"
)

type TravelRequest struct {
	Id uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ApplicantName string `gorm:"column:applicant_name;type:varchar(64);not null"`
	Destination string `gorm:"column:destination;type:varchar(100);not null"`
	DepartureDate time.Time `gorm:"column:departure_date;type:date;not null"`
	ReturnDate time.Time `gorm:"column:return_date;type:date;not null"`
	Status Status `gorm:"type:status; default:'REQUESTED'; column:status"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (t *TravelRequest) SetDates(departureDateStr, returnDateStr string) error {
		const dateFormat = "2006-01-02"

		departureDate, err := time.Parse(dateFormat, departureDateStr)
		if err != nil {
			return fmt.Errorf("an error ocurred to convert departure date: %v", err)
		}
	
		returnDate, err := time.Parse(dateFormat, returnDateStr)
		if err != nil {
			return fmt.Errorf("an error ocurred to convert return date: %v", err)
		}
		t.DepartureDate = departureDate
		t.ReturnDate = returnDate
		return nil
}
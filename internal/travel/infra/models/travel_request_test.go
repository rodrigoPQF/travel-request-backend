package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


func TestSetDatesSuccess(t *testing.T) {
	travelRequest := &TravelRequest{}

	err := travelRequest.SetDates("2023-10-01", "2023-10-01")

	assert.NoError(t, err, "expected success to define date")
	assert.Equal(t, time.Date(2023, 10, 1, 0,0,0,0, time.UTC), travelRequest.DepartureDate)
	assert.Equal(t, time.Date(2023, 10, 1, 0,0,0,0, time.UTC), travelRequest.ReturnDate)
}

func TestSetDatesInvalidDepartureDate(t *testing.T) {
	travelRequest := &TravelRequest{}
	err := travelRequest.SetDates("thisNotDate", "2023-10-10")

	assert.Error(t, err, "expected error to define malformed date")
	assert.Contains(t, err.Error(), "an error ocurred to convert departure date")
}

func TestSetDatesInvalidReturnDate(t *testing.T) {
	travelRequest := &TravelRequest{}
	err := travelRequest.SetDates("2023-10-10", "thisNotDate")

	assert.Error(t, err, "expected error to define malformed date")
	assert.Contains(t, err.Error(), "an error ocurred to convert return date")
}

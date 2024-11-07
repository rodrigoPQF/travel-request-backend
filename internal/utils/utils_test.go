package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateConnectionURL(t *testing.T) {
	tests := []struct {
		name           string
		connectionURL  ConnectionURL
		expectedURL    string
	}{
		{
			name: "Connection URL with credentials and simple password",
			connectionURL: ConnectionURL{
				Username:     "user",
				Password:     "password",
				Host:         "localhost:5432",
				DatabaseName: "dbname",
			},
			expectedURL: "postgres://user:password@localhost:5432/dbname",
		},
		{
			name: "Connection URL with special characters in password",
			connectionURL: ConnectionURL{
				Username:     "user",
				Password:     "p@ssw!rd",
				Host:         "localhost:5432",
				DatabaseName: "dbname",
			},
			expectedURL: "postgres://user:p%40ssw%21rd@localhost:5432/dbname",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateConnectionURL(tt.connectionURL)
			assert.Equal(t, tt.expectedURL, result)
		})
	}
}

func TestValidateStruct(t *testing.T) {
	type ValidationStruct struct {
		Field1 string `validate:"required"`
		Field2 int    `validate:"gte=0"`
	}

	tests := []struct {
		name    string
		input   ValidationStruct
		expectedError bool
	}{
		{
			name: "success validation",
			input: ValidationStruct{
				Field1: "test",
				Field2: 10,
			},
			expectedError: false,
		},
		{
			name: "error validation on field1",
			input: ValidationStruct{
				Field1: "",
				Field2: 10,
			},
			expectedError: true,
		},
		{
			name: "error validation on field2",
			input: ValidationStruct{
				Field1: "test",
				Field2: -1,
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateStruct(tt.input)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
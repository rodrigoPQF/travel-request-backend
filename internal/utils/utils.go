package utils

import (
	"fmt"
	"net/url"

	"github.com/go-playground/validator/v10"
)

type ConnectionURL struct {
	Username string
	Password string
	Host string
	DatabaseName string
}

func GenerateConnectionURL(connectionUrl ConnectionURL) string {
	passwordEscaped := url.QueryEscape(connectionUrl.Password)
	return fmt.Sprintf("postgres://%s:%s@%s/%s",
			connectionUrl.Username,
			passwordEscaped,
			connectionUrl.Host,
			connectionUrl.DatabaseName,
	)
}

var validate = validator.New()

func ValidateStruct(data interface{}) error {
	return validate.Struct(data)
}

type Pagination struct {
	Page int
	Limit int
}

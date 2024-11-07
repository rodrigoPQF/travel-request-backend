package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type BaseResponseHandler struct {
	Status bool `json:"status"`
	Message string `json:"message,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

func NewSuccessResponse(ctx *fiber.Ctx, statusCode int, message string, payload interface{}) error {
	return ctx.Status(statusCode).JSON(BaseResponseHandler{
		Status: true,
		Message: message,
		Payload: payload,
	})
}

func NewErrorResponse(ctx *fiber.Ctx , statusCode int, err string) error {
	return ctx.Status(statusCode).JSON(BaseResponseHandler{
		Status: false,
		Message: err,
	})
}
package controllers

import (
	"authentication/app/dto"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var validate = validator.New()

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func generateHttpError(c *fiber.Ctx, err error, msg string, statusCode int) error {
	zap.L().Error(msg, zap.Error(err), zap.String("path", c.Path()))
	return fiber.NewError(statusCode, msg)
}

func generateHttpInfo(c *fiber.Ctx, msg string) {
	zap.L().Info(msg, zap.String("path", c.Path()))
}

func ValidateStruct[S dto.JwtClaimsInput | dto.ValidateJwtInput](s S) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

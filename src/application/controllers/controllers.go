package controllers

import (
	"time"

	"github.com/JorgeReus/istio-jwt/application/schemas"
	"github.com/JorgeReus/istio-jwt/jwks"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func ValidateMachineClaims(claims schemas.InputJWTClaims) []*schemas.ErrorResponse {
	var errors []*schemas.ErrorResponse
	validate := validator.New()
	err := validate.Struct(claims)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element schemas.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

// GenerateJwk godoc
// @Summary Generates a new Jwk pair
// @Description Randomly generates a jwk pair (public and private)
// @ID generate-jwk
// @Accept  json
// @Produce  json
// @Success 200 {object} schemas.JWKResponse
// @Router /jwk/generate [get]
func GenerateJwk(c *fiber.Ctx) error {
	err := jwks.GenerateKeyPair()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	priv, pub := jwks.GetJson()
	resp := schemas.JWKResponse{
		Public:  pub,
		Private: priv,
	}
	return c.JSON(resp)
}

// GenerateJwt godoc
// @Summary Generates a new JSON Web Token
// @Description Randomly generates a json web token baes on a subject
// @ID generate-jwt
// @Accept  json
// @Produce  json
// @Param b body schemas.InputJWTClaims true "Body for subject"
// @Success 200 {object} string
// @Router /jwt/generate [post]
func GenerateJwt(c *fiber.Ctx) error {
	body := new(schemas.InputJWTClaims)

	if err := c.BodyParser(body); err != nil {
		return err
	}

	errors := ValidateMachineClaims(*body)
	if errors != nil {
		return c.JSON(errors)
	}

	cl := jwt.Claims{
		Subject:  body.Subject,
		Issuer:   "myCompany.com",
		Audience: jwt.Audience{"MyAudience"},
		IssuedAt: jwt.NewNumericDate(time.Now().Local()),
		Expiry:   jwt.NewNumericDate(time.Now().Local().Add(time.Minute * time.Duration(5))),
	}

	jwt := jwks.NewJwt(cl)
	return c.SendString(jwt)
}

// GetPublicJwt godoc
// @Summary Gets the current public JWK
// @Description Gets the public JWK
// @ID get-public-jwk
// @Produce  json
// @Success 200 {object} schemas.PublicJWKResponse
// @Router /jwk/public [get]
func GetPublicJWK(c *fiber.Ctx) error {
	jwk := jwks.PublicJwk()
	return c.JSON(jwk)
}

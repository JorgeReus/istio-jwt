package main

import (
	"log"
	"time"

	_ "github.com/JorgeReus/istio-jwt/docs"
	swagger "github.com/arsmn/fiber-swagger/v2"

	// _ "github.com/arsmn/fiber-swagger/v2/example/docs"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/JorgeReus/istio-jwt/jwks"
	"gopkg.in/square/go-jose.v2/jwt"
)

type TokenResponse struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	IDToken     string `json:"id_token"`
}

type InputJWTClaims struct {
	Subject string `json:"subject" validate:"required,min=3,max=32"`
}

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

type JWKResponse struct {
	Public  string `json:"public_b64"`
	Private string `json:"private_b64"`
}

func ValidateMachineClaims(claims InputJWTClaims) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(claims)
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

func init() {
	err := jwks.GenerateKeyPair()
	if err != nil {
		panic(err)
	}
}

// GenerateJwk godoc
// @Summary Generates a new Jwk pair
// @Description Randomly generates a jwk pair (public and private)
// @ID generate-jwk
// @Accept  json
// @Produce  json
// @Success 200 {object} JWKResponse
// @Router /jwk/generate [get]
func generateJwk(c *fiber.Ctx) error {
	err := jwks.GenerateKeyPair()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	priv, pub := jwks.GetJson()
	resp := JWKResponse{
		Public:  pub,
		Private: priv,
	}
	return c.JSON(resp)

}

func main() {

	app := fiber.New(fiber.Config{})

	app.Use(recover.New())

	app.Get("/swagger/*", swagger.Handler) // default

	app.Get("/jwk/generate", generateJwk)

	app.Post("jwt/generate", func(c *fiber.Ctx) error {
		body := new(InputJWTClaims)

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
	})

	app.Get("/jwk/public", func(c *fiber.Ctx) error {
		jwk := jwks.PublicJwk()
		return c.JSON(jwk)
	})

	log.Fatal(app.Listen(":8080"))
}

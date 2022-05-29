package controllers

import (
	"authentication/app/dto"
	"authentication/core/jwks"

	"github.com/gofiber/fiber/v2"
)

// GenerateJwk godoc
// @Tags jwk
// @Summary Generates a new Jwk pair
// @Description Randomly generates a jwk pair (public and private)
// @ID generate-jwk
// @Accept  json
// @Produce  json
// @Success 201 {object} dto.JwkResponse "JWK pair successfully created"
// @Router /jwk/generate [get]
func GenerateJwk(c *fiber.Ctx) error {
	jwk, err := jwks.New()
	if err != nil {
		return generateHttpError(c, err, "Cannot Generate JWK", fiber.StatusInternalServerError)
	}
	pub := jwk.PublicBase64()
	priv := jwk.PrivateBase64()
	resp := dto.JwkResponse{
		Public:  pub,
		Private: priv,
	}
	return c.Status(fiber.StatusCreated).
		JSON(resp)
}

// PublicJwk godoc
// @Tags jwk
// @Summary Gets the public JWK Set
// @Description Gets the public JWK Set in json format
// @ID public-jwk
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.JwksResponse "JWK public set, for more information refer to https://tools.ietf.org/id/draft-ietf-jose-json-web-key-00.html#:~:text=Abstract,a%20set%20of%20public%20keys."
// @Router /jwk/public [get]
func PublicJwk(c *fiber.Ctx, jwkProvider *jwks.JWK) error {
	resp := jwks.MakePublicJWKS(jwkProvider)
	return c.Status(fiber.StatusOK).
		JSON(resp)
}

package controllers

import (
	"authentication/app/config"
	"authentication/app/dto"
	"authentication/core/jwks"
	"authentication/core/jwt"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	conf = config.GetConfig()
)

// GenerateJwt godoc
// @Tags jwt
// @Summary Generates a new JWT
// @Description Randomly generates a jwt based on claims
// @ID generate-jwt
// @Accept  json
// @Produce  json
// @Success 201 {object} dto.JwtResponse "JWT pair successfully created"
// @Router /jwt/generate [post]
// @Param claims body dto.JwtClaimsInput true "claims body"
func GenerateJwt(c *fiber.Ctx, jwkProvider *jwks.JWK) error {
	body := new(dto.JwtClaimsInput)
	if err := c.BodyParser(body); err != nil {
		return err
	}
	errors := ValidateStruct(*body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	now := time.Now()
	claims := jwt.JwtClaims{
		Issuer:    &conf.JwtIssuer,
		Audience:  &conf.JwtAudience,
		Subject:   &body.Subject,
		JTI:       &body.Subject,
		Name:      &body.FQN,
		Groups:    body.Groups,
		Expiry:    now.Add(time.Minute * 5).Unix(),
		NotBefore: now.Unix(),
		IssuedAt:  now.Unix(),
	}
	err, id_token, refresh_token := jwt.NewJwt(claims, jwkProvider)
	resp := dto.JwtResponse{
		IdToken:      *id_token,
		RefreshToken: *refresh_token,
	}
	if err != nil {
		generateHttpError(c, err, "Cannot Generate Jwt pair", fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusCreated).
		JSON(resp)
}

// ValidateIdToken godoc
// @Tags jwt
// @Summary Validate a new Id token
// @Description Validates an Id token in the form of base64
// @ID validate-id-token
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.JwtValidResponse "Returns wheter the JWT is valid or not"
// @Failure 400 {object} []ErrorResponse "JWT has some error"
// @Router /jwt/validate-id-token [post]
// @Param jwt body dto.ValidateJwtInput true "JWT to validate"
func ValidateIdToken(c *fiber.Ctx, jwkProvider *jwks.JWK) error {
	body := new(dto.ValidateJwtInput)
	if err := c.BodyParser(body); err != nil {
		return err
	}
	errors := ValidateStruct(*body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	resp := dto.JwtValidResponse{
		IsValid: true,
	}
	err := jwt.ValidateIdToken(&body.Jwt, &conf.JwtIssuer, jwkProvider)
	if err != nil {
		resp.IsValid = false
		generateHttpInfo(c, err.Error())
	}
	return c.Status(fiber.StatusCreated).
		JSON(resp)
}

// ValidateRefreshToken godoc
// @Tags jwt
// @Summary Validate a refresh token
// @Description Validates a refresh token in form of a base64 jwt
// @ID validate-refresh-token
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.JwtValidResponse "Returns wheter the JWT is valid or not"
// @Failure 400 {object} []ErrorResponse "JWT has some error"
// @Router /jwt/validate-refresh-token [post]
// @Param jwt body dto.ValidateJwtInput true "JWT to validate"
func ValidateRefreshToken(c *fiber.Ctx, jwkProvider *jwks.JWK) error {
	body := new(dto.ValidateJwtInput)
	if err := c.BodyParser(body); err != nil {
		return err
	}
	errors := ValidateStruct(*body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	resp := dto.JwtValidResponse{
		IsValid: true,
	}
	err := jwt.ValidateRefreshToken(&body.Jwt, jwkProvider)
	if err != nil {
		resp.IsValid = false
		generateHttpInfo(c, err.Error())
	}
	return c.Status(fiber.StatusCreated).
		JSON(resp)
}

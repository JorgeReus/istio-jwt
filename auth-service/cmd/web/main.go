package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "authentication/app/_docs"
	"authentication/app/config"
	"authentication/app/controllers"
	"authentication/core/jwks"

	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// @title JWT authentication API
// @version 0.0.1
// @description Sample service for showcasing istio capabilities in edge JWT authentication
func main() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()
	fiberConfig := fiber.Config{}
	app := fiber.New(fiberConfig)
	app.Use(recover.New())

	appConfig := config.GetConfig()
	if disable_swagger := appConfig.DisableSwagger; disable_swagger != "true" {
		config := swagger.Config{
			URL:         fmt.Sprintf("%s/swagger/doc.json", config.GetConfig().SchemeDomainName),
			DeepLinking: true,
		}
		app.Get("/swagger/*", swagger.New(config))
	}

	err, jwkProvider := jwks.LoadFromBase64(&appConfig.PublicJwk, &appConfig.PrivateJwk)
	if err != nil {
		log.Fatal(err)
	}

	jwkGroup := app.Group("/jwk")
	jwkGroup.Get("/generate", controllers.GenerateJwk)
	jwkGroup.Get("/public", func(c *fiber.Ctx) error {
		return controllers.PublicJwk(c, jwkProvider)
	})

	jwtGroup := app.Group("/jwt")
	jwtGroup.Post("/generate", func(c *fiber.Ctx) error {
		return controllers.GenerateJwt(c, jwkProvider)
	})
	jwtGroup.Post("/validate-id-token", func(c *fiber.Ctx) error {
		return controllers.ValidateIdToken(c, jwkProvider)
	})
	jwtGroup.Post("/validate-refresh-token", func(c *fiber.Ctx) error {
		return controllers.ValidateRefreshToken(c, jwkProvider)
	})

	app.Listen(":8080")
}

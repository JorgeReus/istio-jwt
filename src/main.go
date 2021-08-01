package main

import (
	"log"

	"github.com/JorgeReus/istio-jwt/application/controllers"

	_ "github.com/JorgeReus/istio-jwt/docs"
	"github.com/JorgeReus/istio-jwt/jwks"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "go.uber.org/automaxprocs"
)

func init() {
	err := jwks.GenerateKeyPair()
	if err != nil {
		panic(err)
	}
}

func main() {

	app := fiber.New(fiber.Config{})

	app.Use(recover.New())

	app.Get("/swagger/*", swagger.Handler) // default

	app.Get("/jwk/generate", controllers.GenerateJwk)

	app.Post("/jwt/generate", controllers.GenerateJwt)

	app.Get("/jwk/public", controllers.GetPublicJWK)

	app.Get("/healthz", controllers.Healthz)

	app.Get("/readiness", controllers.ReadinessProbe)

	log.Fatal(app.Listen(":8080"))
}

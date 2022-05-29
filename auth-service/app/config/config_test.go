package config

import (
	"log"
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func cleanEnv() {
	os.Clearenv()
}

func setRequiredEnvs() {
	os.Setenv("PUBLIC_JWK", "asd")
	os.Setenv("PRIVATE_JWK", "asd")
	os.Setenv("JWT_ISSUER", "a")
	os.Setenv("JWT_AUDIENCE", "a")
	os.Setenv("SCHEME_DOMAIN_NAME", "asd")
}

func testRequired(t *testing.T) {
	fakeLogFatal := func(msg ...interface{}) {
		panic("log.Fatal called")
	}
	patch := monkey.Patch(log.Fatal, fakeLogFatal)
	defer patch.Unpatch()
	setRequiredEnvs()
	GetConfig()
}

func testRequiredFail(t *testing.T) {
	c = nil
	fakeLogFatal := func(msg ...interface{}) {
		panic("log.Fatal called")
	}
	patch := monkey.Patch(log.Fatal, fakeLogFatal)
	defer patch.Unpatch()
	assert.Panics(t, func() { GetConfig() }, "Not panic")
}

func TestController(t *testing.T) {
	fs := map[string]func(*testing.T){"testRequired": testRequired, "testRequiredFail": testRequiredFail}
	for name, f := range fs {
		cleanEnv()
		t.Run(name, f)
	}
}

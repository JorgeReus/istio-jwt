package main

import (
	"authentication/core/jwks"
	"fmt"
	"os"
	"path/filepath"
)

type EnvFile struct {
	PrivJwk          string `yaml:"PRIVATE_JWK"`
	PubJwk           string `yaml:"PUBLIC_JWK"`
	Issuer           string `yaml:"JWT_ISSUER"`
	Audience         string `yaml:"JWT_AUDIENCE"`
	SchemeDomainName string `yaml:"SCHEME_DOMAIN_NAME"`
}

func main() {
	jwks, err := jwks.New()
	if err != nil {
		panic(err)
	}

	newpath := filepath.Join(".", "out", jwks.Private().KeyID)
	err = os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	// Public
	writeFile(fmt.Sprintf("%s/public.b64.txt", newpath), jwks.PublicBase64())
	pubKey, err := jwks.Public().MarshalJSON()
	if err != nil {
		panic(err)
	}
	writeFile(fmt.Sprintf("%s/public.json", newpath), string(pubKey))

	// Private
	writeFile(fmt.Sprintf("%s/private.b64.txt", newpath), jwks.PrivateBase64())
	privKey, err := jwks.Private().MarshalJSON()
	if err != nil {
		panic(err)
	}
	writeFile(fmt.Sprintf("%s/private.json", newpath), string(privKey))

	var envData map[string]string = map[string]string{
		"PRIVATE_JWK":        jwks.PrivateBase64(),
		"PUBLIC_JWK":         jwks.PublicBase64(),
		"JWT_ISSUER":         "my.jwt.issuer",
		"JWT_AUDIENCE":       "my.jwt.audience",
		"SCHEME_DOMAIN_NAME": "http://localhost:8080",
	}

	var envString string

	for k, v := range envData {
		envString += fmt.Sprintln(fmt.Sprintf("%s=%s", k, v))
	}

	writeFile(".env", envString)
}

func writeFile(path, contents string) {
	public, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer public.Close()
	public.WriteString(contents)
}

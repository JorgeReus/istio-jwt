package jwks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateKeys(t *testing.T) *JWK {
	jwkPair, err := New()
	assert.NoError(t, err)
	assert.NotNil(t, jwkPair)
	return jwkPair
}

func TestGenerate(t *testing.T) {
	jwkPair := generateKeys(t)
	priv := jwkPair.Private()
	public := jwkPair.Public()
	assert.NotNil(t, priv)
	assert.NotNil(t, public)
}

func TestLoadFromBase64(t *testing.T) {
	jwkPair := generateKeys(t)
	public := jwkPair.PublicBase64()
	private := jwkPair.PrivateBase64()
	err, jwk := LoadFromBase64(&public, &private)
	assert.NoError(t, err)
	assert.NotNil(t, jwk)
	t.Run("testInvalidBase64", func(t *testing.T) {
		pub := "asd"
		priv := "asd"
		err, _ := LoadFromBase64(&pub, &priv)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Couldn't decode base64 public JWK")

		err, _ = LoadFromBase64(&public, &priv)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Couldn't decode base64 private JWK")
	})

	t.Run("testInvalidJwk", func(t *testing.T) {
		pub := "dGVzdAo="
		priv := "dGVzdAo="
		err, _ := LoadFromBase64(&pub, &priv)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Couldn't parse public JWK")

		err, _ = LoadFromBase64(&public, &priv)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Couldn't parse private JWK")
	})
}

func TestMakeSet(t *testing.T) {
	jwkPair := generateKeys(t)
	privSet := MakePrivateJWKS(jwkPair)
	assert.NotEmpty(t, privSet)
	pubSet := MakePublicJWKS(jwkPair)
	assert.NotEmpty(t, pubSet)
}

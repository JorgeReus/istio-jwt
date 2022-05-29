package jwt

import (
	"authentication/core/jwks"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	key    *jwks.JWK
	issuer = "ceneval.edu.mx"
)

func init() {
	var err error
	key, err = jwks.New()
	if err != nil {
		log.Fatalln(err)
	}
}

func TestInvalidJwtIssuer(t *testing.T) {
	_, userInfoTok, refreshTok := createValidJwtPair(issuer)

	randomIssuer := "issuer"
	err := ValidateIdToken(userInfoTok, &randomIssuer, key)
	assert.EqualError(t, err, "Error Validating JWT, issuer didn't match")

	err = ValidateRefreshToken(refreshTok, key)
}

func TestInvalidClaimsJwt(t *testing.T) {
	raw := "eyJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJpc3N1ZXIiLCJzdWIiOiJzdWJqZWN0In0.gpHyA1B1H6X4a4Edm9wo7D3X2v3aLSDBDG2_5BzXYe0"
	randomIssuer := "ceneval"
	err := ValidateIdToken(&raw, &randomIssuer, key)
	assert.EqualError(t, err, "Invalid JWT :square/go-jose: error in cryptographic primitive")
}

func TestInvalidJwt(t *testing.T) {
	raw := "asdasdasd"
	randomIssuer := "ceneval"
	err := ValidateIdToken(&raw, &randomIssuer, key)
	assert.EqualError(t, err, "Couldn't parse base64 jwt")
}

func createValidJwtPair(issuer string) (error, *string, *string) {
	audience := "internal"
	subject := "subject"

	err, userInfoTok, refreshTok := NewJwt(JwtClaims{
		Issuer:   &issuer,
		Audience: &audience,
		Subject:  &subject,
		Groups:   []string{},
		Expiry:   time.Now().Add(time.Minute).Unix(),
	}, key)
	return err, userInfoTok, refreshTok
}

func TestValidJwt(t *testing.T) {
	err, userInfoTok, refreshTok := createValidJwtPair(issuer)
	assert.Nil(t, err)
	assert.NotEmpty(t, userInfoTok)
	assert.NotEmpty(t, refreshTok)
	err = ValidateIdToken(userInfoTok, &issuer, key)
	assert.Nil(t, err)
}

func TestJwtJwkNil(t *testing.T) {
	_, userInfoTok, _ := createValidJwtPair(issuer)
	err := ValidateIdToken(userInfoTok, &issuer, nil)
	assert.EqualError(t, err, "Error, creating JWT, key must not be nil")
}

func TestJwtJwkEmpty(t *testing.T) {
	_, userInfoTok, _ := createValidJwtPair(issuer)
	err := ValidateIdToken(userInfoTok, &issuer)
	assert.EqualError(t, err, "Error creating JWT, you need to supply at least 1 JWK")
}

func TestErrorJwkNil(t *testing.T) {
	issuer := "ceneval.edu.mx"
	audience := "internal"
	subject := "subject"

	err, _, _ := NewJwt(JwtClaims{
		Issuer:   &issuer,
		Audience: &audience,
		Subject:  &subject,
		Groups:   []string{},
		Expiry:   time.Now().Add(time.Minute).Unix(),
	}, nil)

	assert.EqualError(t, err, "Error, creating JWT, key must not be nil")
}

func TestErrorJwkNotPresent(t *testing.T) {
	issuer := "ceneval.edu.mx"
	audience := "internal"
	subject := "subject"

	err, _, _ := NewJwt(JwtClaims{
		Issuer:   &issuer,
		Audience: &audience,
		Subject:  &subject,
		Groups:   []string{},
		Expiry:   time.Now().Add(time.Minute).Unix(),
	})

	assert.EqualError(t, err, "Error creating JWT, you need to supply at least 1 JWK")
}

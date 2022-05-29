package jwt

import (
	"authentication/core/jwks"
	"errors"
	"fmt"
	"time"

	math_rand "math/rand"

	jose "gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

type JwtClaims struct {
	JTI       *string
	Issuer    *string
	Audience  *string
	Subject   *string
	Name      *string
	Groups    []string
	NotBefore int64
	Expiry    int64
	IssuedAt  int64
}

func NewJwt(claims JwtClaims, keys ...*jwks.JWK) (error, *string, *string) {
	if len(keys) < 1 {
		return errors.New("Error creating JWT, you need to supply at least 1 JWK"), nil, nil
	}
	headers := make(map[jose.HeaderKey]interface{})
	jwk := keys[math_rand.Intn(len(keys))]
	if jwk == nil {
		return errors.New("Error, creating JWT, key must not be nil"), nil, nil
	}
	headers[jose.HeaderKey("kid")] = jwk.Private().KeyID
	opts := jose.SignerOptions{ExtraHeaders: headers}
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jwks.Alg, Key: jwk.Private().Key}, &opts)
	if err != nil {
		return err, nil, nil
	}

	raw, err := jwt.Signed(signer).Claims(map[string]interface{}{
		"iss":    claims.Issuer,
		"aud":    claims.Audience,
		"sub":    claims.Subject,
		"jti":    claims.JTI,
		"name":   claims.Name,
		"groups": claims.Groups,
		"exp":    claims.Expiry,
		"nbf":    claims.NotBefore,
		"iat":    claims.IssuedAt,
	}).CompactSerialize()
	if err != nil {
		return err, nil, nil
	}

	refreshToken, err := jwt.Signed(signer).Claims(map[string]interface{}{
		"user_id": claims.Subject,
	}).CompactSerialize()

	return nil, &raw, &refreshToken
}

func ValidateIdToken(jwtBase64 *string, issuer *string, keys ...*jwks.JWK) error {
	if len(keys) < 1 {
		return errors.New("Error creating JWT, you need to supply at least 1 JWK")
	}
	jwk := keys[math_rand.Intn(len(keys))]
	if jwk == nil {
		return errors.New("Error, creating JWT, key must not be nil")
	}
	token, err := jwt.ParseSigned(*jwtBase64)
	if err != nil {
		return errors.New("Couldn't parse base64 jwt")
	}

	claims := jwt.Claims{}

	// Validate if the JWT was created with the supplied JWK & map the claims
	if err := token.Claims(jwk.Public(), &claims); err != nil {
		return errors.New(fmt.Sprintf("Invalid JWT :%v", err))
	}

	// Validate the exp time if supplied
	if claims.Expiry.Time().Before(time.Now()) {
		return errors.New(fmt.Sprintf("Id token expired at %s", claims.Expiry.Time().String()))
	}

	// Validate if the issuer is ok
	err = claims.Validate(jwt.Expected{
		Issuer: *issuer,
	})

	if err != nil {
		return errors.New("Error Validating JWT, issuer didn't match")
	}
	return nil
}

func ValidateRefreshToken(jwtBase64 *string, keys ...*jwks.JWK) error {
	if len(keys) < 1 {
		return errors.New("Error creating JWT, you need to supply at least 1 JWK")
	}
	jwk := keys[math_rand.Intn(len(keys))]
	if jwk == nil {
		return errors.New("Error, creating JWT, key must not be nil")
	}
	token, err := jwt.ParseSigned(*jwtBase64)
	if err != nil {
		return errors.New("Couldn't parse base64 jwt")
	}

	claims := map[string]interface{}{}

	// Validate if the JWT was created with the supplied JWK & map the claims
	if err := token.Claims(jwk.Public(), &claims); err != nil {
		return errors.New(fmt.Sprintf("Invalid JWT :%v", err))
	}
	if _, ok := claims["user_id"]; !ok && len(claims) != 1 {
		return errors.New("Invalid Id token, claims do not match")
	}
	return nil
}

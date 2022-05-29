package jwks

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	jose "gopkg.in/square/go-jose.v2"
)

type JWK struct {
	public        jose.JSONWebKey
	private       jose.JSONWebKey
	base64Public  string
	base64Private string
}

var (
	Alg   = jose.ES256
	curve = elliptic.P256()
)

func New() (jwk *JWK, err error) {
	use := "sig"

	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}

	hasher := sha256.New()
	hasher.Write(privKey.X.Bytes())
	kid := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	privateJwk := jose.JSONWebKey{Key: privKey, KeyID: kid, Algorithm: string(Alg), Use: use}
	publicJwk := privateJwk.Public()

	privateRaw, err := privateJwk.MarshalJSON()
	if err != nil {
		return nil, err
	}

	publicRaw, err := publicJwk.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return &JWK{
		public:        publicJwk,
		private:       privateJwk,
		base64Public:  base64.URLEncoding.EncodeToString(publicRaw),
		base64Private: base64.URLEncoding.EncodeToString(privateRaw),
	}, nil
}

func LoadFromBase64(publicB64 *string, privateB64 *string) (error, *JWK) {
	publicData, err := base64.URLEncoding.DecodeString(*publicB64)
	if err != nil {
		return errors.New(fmt.Sprintf("Couldn't decode base64 public JWK: %v", err)), nil
	}
	privateData, err := base64.URLEncoding.DecodeString(*privateB64)
	if err != nil {
		return errors.New(fmt.Sprintf("Couldn't decode base64 private JWK: %v", err)), nil
	}
	public := jose.JSONWebKey{}
	err = json.Unmarshal(publicData, &public)
	if err != nil {
		return errors.New(fmt.Sprintf("Couldn't parse public JWK: %v", err)), nil
	}

	private := jose.JSONWebKey{}
	err = json.Unmarshal(privateData, &private)

	if err != nil {
		return errors.New(fmt.Sprintf("Couldn't parse private JWK: %v", err)), nil
	}
	return nil, &JWK{
		private:       private,
		public:        public,
		base64Private: *privateB64,
		base64Public:  *publicB64,
	}
}

func MakePublicJWKS(jwks ...*JWK) jose.JSONWebKeySet {
	keys := []jose.JSONWebKey{}
	for _, jwk := range jwks {
		keys = append(keys, jwk.public)
	}
	return jose.JSONWebKeySet{Keys: keys}
}

func MakePrivateJWKS(jwks ...*JWK) jose.JSONWebKeySet {
	keys := []jose.JSONWebKey{}
	for _, jwk := range jwks {
		keys = append(keys, jwk.private)
	}
	return jose.JSONWebKeySet{Keys: keys}
}

func (jwk *JWK) Public() jose.JSONWebKey {
	return jwk.public
}

func (jwk *JWK) Private() jose.JSONWebKey {
	return jwk.private
}

func (jwk *JWK) PublicBase64() string {
	return jwk.base64Public
}

func (jwk *JWK) PrivateBase64() string {
	return jwk.base64Private
}

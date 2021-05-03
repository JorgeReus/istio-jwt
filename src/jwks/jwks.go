package jwks

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"

	jose "gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

var publicJwk jose.JSONWebKey
var privateJwk jose.JSONWebKey

var publicJsonJwk string
var privateJsonJwk string

func NewJwt(claims jwt.Claims) string {
	// fmt.Println(ExportRsaPrivateKeyAsPemStr(jwkNew.Key.(*rsa.PrivateKey)))
	// fmt.Println(ExportRsaPublicKeyAsPemStr(jwkNew.Public().Key.(*rsa.PublicKey)))

	headers := make(map[jose.HeaderKey]interface{})
	headers[jose.HeaderKey("kid")] = privateJwk.KeyID
	opts := jose.SignerOptions{ExtraHeaders: headers}
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: privateJwk.Key}, &opts)
	if err != nil {
		panic(err)
	}

	raw, err := jwt.Signed(signer).Claims(claims).CompactSerialize()
	if err != nil {
		panic(err)
	}

	return raw
}

func GenerateKeyPair() (err error) {

	keySize := 2048
	use := "sig"
	alg := string(jose.RS256)

	privKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return err
	}

	hasher := sha256.New()
	hasher.Write(privKey.N.Bytes())
	kid := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	privateJwk = jose.JSONWebKey{Key: privKey, KeyID: kid, Algorithm: alg, Use: use}
	publicJwk = privateJwk.Public()

	privateRaw, err := privateJwk.MarshalJSON()
	if err != nil {
		return err
	}

	publicRaw, err := publicJwk.MarshalJSON()
	if err != nil {
		return err
	}

	privateJsonJwk = base64.URLEncoding.EncodeToString(privateRaw)
	publicJsonJwk = base64.URLEncoding.EncodeToString(publicRaw)

	return nil
}

func ExportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) []byte {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	return pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
}

func PublicJwk() jose.JSONWebKeySet {
	set := jose.JSONWebKeySet{Keys: []jose.JSONWebKey{publicJwk}}
	return set
}

func GetJson() (string, string) {
	return privateJsonJwk, publicJsonJwk
}

func ExportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)

	return string(pubkey_pem), nil
}

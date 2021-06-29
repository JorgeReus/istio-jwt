package schemas

type TokenResponse struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	IDToken     string `json:"id_token"`
}

type InputJWTClaims struct {
	Subject string `json:"subject" validate:"required,min=3,max=32"`
}

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

type JWKResponse struct {
	Public  string `json:"public_b64"`
	Private string `json:"private_b64"`
}

type PublicJWKResponse struct {
	Keys []struct {
		Use string `json:"use"`
		Kty string `json:"kty"`
		Kid string `json:"kid"`
		Alg string `json:"alg"`
		N   string `json:"n"`
		E   string `json:"e"`
	} `json:"keys"`
}

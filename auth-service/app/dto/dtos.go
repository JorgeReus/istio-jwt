package dto

type JwtClaimsInput struct {
	Subject string   `json:"subject" example:"Subject of the JWT" validate:"required,min=3,max=32"`
	FQN     string   `json:"fqn" example:"Fully Qualified Name of the Subject" validate:"required,min=3,max=40"`
	Groups  []string `json:"groups" example:"group 1,group 2" validate:"required,min=1,max=10"`
}

type ValidateJwtInput struct {
	Jwt string `json:"jwt" example:"Jwt in base64 form" validate:"required,min=10"`
}

type JWK struct {
	Use                string `json:"use" example:"sig"`
	KeyType            string `json:"kty" example:"EC"`
	KeyID              string `json:"kid" example:"a2V5SWRGb3JUZXN0aW5nCg=="`
	CryptographicCurve string `json:"crv" example:"P-256"`
	Algorithm          string `json:"alg" example:"ES256"`
	X                  string `json:"x" example:"eEVsaXB0aWNDb29yZGluYXRlRm9ydGVzdGluZwo="`
	Y                  string `json:"y" example:"eUVsaXB0aWNDb29yZGluYXRlRm9ydGVzdGluZwo="`
}

type JwksResponse []JWK

type JwtResponse struct {
	IdToken      string `json:"id_token" example:"eyJhbGciOiJFUzI1NiIsImtpZCI6IlN6Qy1Ta2VUMl9vYTFwZ3lzOXpGV2VmM3hKTXZTUmtpdlBqQkloeDQ4cEk9In0.eyJhdWQiOiJpbnRlcm5hbC1jZW5ldmFsLWVtcGxveWVlcyIsImV4cCI6MTY0OTk5MjQ1OSwiZ3JvdXBzIjpudWxsLCJpYXQiOjE2NDk5ODg4NTksImlzcyI6InBsYXRhZm9ybWEuY2VuZXZhbC5lZHUubXgiLCJzdWIiOiJzdHJpbmcifQ.GCz1NTtAZyQMXYj9--dVXOn1k-0nezqP4TkP_WtI3Be4r6SCiuExmMIc6Wlb-tE48oaLaxEylJCxVimK-BdOaw"`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJFUzI1NiIsImtpZCI6IlN6Qy1Ta2VUMl9vYTFwZ3lzOXpGV2VmM3hKTXZTUmtpdlBqQkloeDQ4cEk9In0.eyJ1c2VyX2lkIjoic3RyaW5nIn0.5zL9n6ZMtmeTH4gEKKOiymPkIJYZdDy5Xq_S621PHcEQVjuYsKSYlmZBmIvRd6NfltBQbUR-VPzRmFwIgYKQvA"`
}

type JwkResponse struct {
	Public  string `json:"public_b64" example:"eyJ1c2UiOiJzaWciLCJrdHkiOiJFQyIsImtpZCI6IlZMRmZNck9VbVpOM1lISThrc0xkeFktUVJoX2lFd2dscVh4N2I1TUVBeHM9IiwiY3J2IjoiUC0yNTYiLCJhbGciOiJFUzI1NiIsIngiOiJiakR5c3VfZ09McXZPSFF4cnJwa3YxVDhOSFR6aUVUU0NySHpWM3luZ0djIiwieSI6IlV5bTZEWlRObHBIaDZvRUl4YkRDb0Jya05Lb01DU3ZQOW9GOVJtNFpVdjgifQ=="`
	Private string `json:"private_b64" example:"eyJ1c2UiOiJzaWciLCJrdHkiOiJFQyIsImtpZCI6IlZMRmZNck9VbVpOM1lISThrc0xkeFktUVJoX2lFd2dscVh4N2I1TUVBeHM9IiwiY3J2IjoiUC0yNTYiLCJhbGciOiJFUzI1NiIsIngiOiJiakR5c3VfZ09McXZPSFF4cnJwa3YxVDhOSFR6aUVUU0NySHpWM3luZ0djIiwieSI6IlV5bTZEWlRObHBIaDZvRUl4YkRDb0Jya05Lb01DU3ZQOW9GOVJtNFpVdjgiLCJkIjoiVkJDYkpkang1ajFxWUg4ek1JYUNJOVpXRWVac2JlUWxIc1pZQzRJMnRnNCJ9"`
}

type JwtValidResponse struct {
	IsValid bool `json:"is_valid" example:"false"`
}

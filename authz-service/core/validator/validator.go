package validator

import (
	"github.com/golang-jwt/jwt/v4"
)

func validate(roles []interface{}) bool {
	for _, role := range roles {
		if role.(string) == "admin" {
			return true
		}
	}
	return false

}

var parser = &jwt.Parser{UseJSONNumber: true}

func IsJWTAuthorized(rawToken *string) bool {
	token, _, _ := parser.ParseUnverified(*rawToken, jwt.MapClaims{})
	claims := token.Claims.(jwt.MapClaims)
	roles := claims["groups"].([]interface{})
	if validate(roles) {
		return true
	}
	return false
}

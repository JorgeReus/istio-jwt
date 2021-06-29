// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/jwk/generate": {
            "get": {
                "description": "Randomly generates a jwk pair (public and private)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Generates a new Jwk pair",
                "operationId": "generate-jwk",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.JWKResponse"
                        }
                    }
                }
            }
        },
        "/jwk/public": {
            "get": {
                "description": "Gets the public JWK",
                "produces": [
                    "application/json"
                ],
                "summary": "Gets the current public JWK",
                "operationId": "get-public-jwk",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.PublicJWKResponse"
                        }
                    }
                }
            }
        },
        "/jwt/generate": {
            "post": {
                "description": "Randomly generates a json web token baes on a subject",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Generates a new JSON Web Token",
                "operationId": "generate-jwt",
                "parameters": [
                    {
                        "description": "Body for subject",
                        "name": "b",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.InputJWTClaims"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schemas.InputJWTClaims": {
            "type": "object",
            "required": [
                "subject"
            ],
            "properties": {
                "subject": {
                    "type": "string"
                }
            }
        },
        "schemas.JWKResponse": {
            "type": "object",
            "properties": {
                "private_b64": {
                    "type": "string"
                },
                "public_b64": {
                    "type": "string"
                }
            }
        },
        "schemas.PublicJWKResponse": {
            "type": "object",
            "properties": {
                "keys": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "alg": {
                                "type": "string"
                            },
                            "e": {
                                "type": "string"
                            },
                            "kid": {
                                "type": "string"
                            },
                            "kty": {
                                "type": "string"
                            },
                            "n": {
                                "type": "string"
                            },
                            "use": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
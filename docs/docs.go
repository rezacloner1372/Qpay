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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/signup": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Register Users in Qpay",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Register Users",
                "operationId": "signup",
                "parameters": [
                    {
                        "description": "user first name",
                        "name": "name",
                        "in": "body",
                        "required": true
                    },
                    {
                        "description": "user last name",
                        "name": "family",
                        "in": "body",
                        "required": true
                    },
                    {
                        "description": "user email",
                        "name": "email",
                        "in": "body",
                        "required": true
                    },
                    {
                        "description": "user cellphone",
                        "name": "cellphone",
                        "in": "body",
                        "required": true
                    },
                    {
                        "description": "prefered username",
                        "name": "username",
                        "in": "body",
                        "required": true
                    },
                    {
                        "description": "strong password",
                        "name": "password",
                        "in": "body",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Login users in Qpay",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Login Users",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "user account username",
                        "name": "username",
                        "in": "body",
                        "required": true
                    },
                    {
                        "description": "user account password",
                        "name": "password",
                        "in": "body",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/auth/logout": {
            "get": {
                "description": "Used to get users loged out",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Logout Users",
                "operationId": "logout",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/payment/gateway/new": {
            "post": {
                "description": "Used By users to create new personal payment gateway.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payment Gateway"
                ],
                "summary": "Create Personal Payment Gateway",
                "operationId": "personal-payment-gateway",
                "parameters": [
                    {

                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/payment/gateway/business/new": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Used By Companies to create new business payment gateway.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payment Gateway"
                ],
                "summary": "Create Business Payment Gateway",
                "operationId": "business-payment-gateway",
                "parameters": [

                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/tariff/new": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Define New Tariff for User Payment Gateways.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tariff"
                ],
                "summary": "Create Tariff",
                "operationId": "create-tariff",
                "parameters": [
                    {
                        "type": "string",
                        "description": "prefered name for tariff",
                        "name": "name",
                        "in": "body",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "description for this tariff",
                        "name": "description",
                        "in": "body",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tariff price",
                        "name": "price",
                        "in": "body",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "prefered currency for tariff",
                        "name": "currency",
                        "in": "body",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "validity days of tariff",
                        "name": "validitydays",
                        "in": "body",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "default value can be 0 or not",
                        "name": "isdefault",
                        "in": "body",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/tariff/:id": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get an specific tariff using its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tariff"
                ],
                "summary": "Get Tariff",
                "operationId": "get-tariff",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                },
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update prefered tariff values",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tariff"
                ],
                "summary": "Update Tariff",
                "operationId": "update-tariff",
                "parameters": [
                    {
                        "type": "string",
                        "description": "prefered name for tariff",
                        "name": "name",
                        "in": "body",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "description for this tariff",
                        "name": "description",
                        "in": "body",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tariff price",
                        "name": "price",
                        "in": "body",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "prefered currency for tariff",
                        "name": "currency",
                        "in": "body",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "validity days of tariff",
                        "name": "validitydays",
                        "in": "body",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "default value can be 0 or not",
                        "name": "isdefault",
                        "in": "body",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                },
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete prefered tariff",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tariff"
                ],
                "summary": "Delete Tariff",
                "operationId": "delete-tariff",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                },
            }
        },
        "/tariff/all": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get list of all tariffs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tariff"
                ],
                "summary": "Get All Tariffs",
                "operationId": "get-all-tariffs",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
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
	Version:     "1.0",
	Host:        "127.0.0.1:8080",
	BasePath:    "/",
	Schemes:     []string{"http", "https"},
	Title:       "QPay API",
	Description: "QPay API",
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

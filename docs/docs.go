// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/electromart/v1/customers": {
            "get": {
                "security": [
                    {
                        "AdminAuth": []
                    }
                ],
                "description": "Get all customers",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Get all customers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/electromart/v1/customers/": {
            "get": {
                "security": [
                    {
                        "AdminAuth": []
                    }
                ],
                "description": "Get all customers",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Get all customers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "utils.ErrorResponse": {
            "description": "Error Response with message and statuscode",
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "statuscode": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Electromart API",
	Description:      "This is a json RESTful API for the newly established e-commerce Electromart",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

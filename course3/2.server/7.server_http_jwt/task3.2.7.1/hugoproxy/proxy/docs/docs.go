// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/address/geocode": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Post Address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "list"
                ],
                "summary": "Search",
                "parameters": [
                    {
                        "description": "request",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/address.GeocodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/address.GeocodeResponse"
                        }
                    },
                    "400": {
                        "description": "Неверный формат запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "404 not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Cервис https://dadata.ru не доступен",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/address/search": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Post Address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "list"
                ],
                "summary": "Search",
                "parameters": [
                    {
                        "description": "request",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/address.SearchRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/address.SearchResponse"
                        }
                    },
                    "400": {
                        "description": "Неверный формат запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "404 not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Cервис https://dadata.ru не доступен",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "login account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "request",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.LoginResponse"
                        }
                    },
                    "403": {
                        "description": "invalid token",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/register": {
            "post": {
                "description": "Post Register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register",
                "operationId": "create-account",
                "parameters": [
                    {
                        "description": "request",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.RegisterResponse"
                        }
                    },
                    "403": {
                        "description": "invalid token",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "address.Address": {
            "type": "object",
            "properties": {
                "lat": {
                    "type": "string"
                },
                "lon": {
                    "type": "string"
                }
            }
        },
        "address.GeocodeRequest": {
            "type": "object",
            "properties": {
                "lat": {
                    "type": "string"
                },
                "lon": {
                    "type": "string"
                }
            }
        },
        "address.GeocodeResponse": {
            "type": "object",
            "properties": {
                "addresses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/address.Address"
                    }
                }
            }
        },
        "address.SearchRequest": {
            "type": "object",
            "properties": {
                "query": {
                    "type": "string"
                }
            }
        },
        "address.SearchResponse": {
            "type": "object",
            "properties": {
                "addresses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/address.Address"
                    }
                }
            }
        },
        "main.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "main.LoginResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "main.RegisterRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "main.RegisterResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "Login"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "HugoMap",
	Description:      "API Server for HugoMap Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "DIT - IFAL",
            "email": "evs10@aluno.ifal.edu.br"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Rota que permite que um usuário se autentique no sistema utilizando seu endereço de e-mail e senha.\n| E-mail              | Senha     | Função                                                            |\n|---------------------|-----------|-------------------------------------------------------------------|\n| admin@ifal.edu.br   | Test1234! | Usuário comum do sistema\t\t\t\t\t\t\t\t\t\t  |",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rotas de autenticação"
                ],
                "summary": "Fazer login no sistema",
                "operationId": "Login",
                "responses": {
                    "400": {
                        "description": "Requisição mal formulada.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "401": {
                        "description": "Usuário não autorizado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "403": {
                        "description": "Acesso negado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "422": {
                        "description": "Algum dado informado não pôde ser processado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Ocorreu um erro inesperado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "503": {
                        "description": "A base de dados não está disponível.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.ErrorMessage": {
            "type": "object",
            "properties": {
                "duplicated_fields": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "invalid_fields": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.InvalidField"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "response.InvalidField": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "field_name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "SPOTIFY RECORDS API",
	Description:      "Aplicação de artistas do spotify",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
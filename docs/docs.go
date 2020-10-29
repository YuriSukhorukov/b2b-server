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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth": {
            "post": {
                "description": "Возвращает результат операции авторизации HttpOnly Cookie JWT пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Авторизация пользователя проверкой HttpOnly Cookie JWT",
                "responses": {
                    "200": {
                        "description": "Успешное выполнение операции",
                        "schema": {
                            "$ref": "#/definitions/model.Success"
                        }
                    },
                    "400": {
                        "description": "Неудачная авторизация HttpOnly Cookie JWT",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "401": {
                        "description": "Токен JWT отсутствует",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/email_free/{email}": {
            "get": {
                "description": "Возвращает результат проверки доступности e-mail для регистрации",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Возвращает результат проверки доступности e-mail для регистрации",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Email свободен для регистрации",
                        "schema": {
                            "$ref": "#/definitions/model.Success"
                        }
                    },
                    "400": {
                        "description": "Email занят для регистрации",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/offers": {
            "get": {
                "tags": [
                    "offers"
                ]
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "offers"
                ],
                "parameters": [
                    {
                        "description": "Offer",
                        "name": "offer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Offer"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Успешное выполнение операции",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Record"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/offers/{id}": {
            "get": {
                "tags": [
                    "offers"
                ]
            },
            "delete": {
                "tags": [
                    "offers"
                ]
            },
            "patch": {
                "tags": [
                    "offers"
                ]
            }
        },
        "/signin": {
            "post": {
                "description": "Возвращает результат операции создания HttpOnly Cookie JWT пользователя при авторизации",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Добавляет HttpOnly Cookie JWT пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное выполнение операции",
                        "schema": {
                            "$ref": "#/definitions/model.Record"
                        }
                    },
                    "400": {
                        "description": "Неверный Email или Password",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/signout": {
            "delete": {
                "description": "Возвращает результат операции удаления HttpOnly Cookie JWT пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Удаляет HttpOnly Cookie JWT пользователя",
                "responses": {
                    "200": {
                        "description": "Успешное выполнение операции",
                        "schema": {
                            "$ref": "#/definitions/model.Success"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "Возвращает результат операции добавленя нового пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Добавляет нового пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Успешное выполнение операции",
                        "schema": {
                            "$ref": "#/definitions/model.Record"
                        }
                    },
                    "400": {
                        "description": "Email занят для регистрации",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "something wrong"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "model.Offer": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string",
                    "format": "int",
                    "example": "100"
                },
                "city": {
                    "type": "string",
                    "format": "string",
                    "example": "Москва"
                },
                "country": {
                    "type": "string",
                    "format": "string",
                    "example": "Российская Федерация"
                },
                "created_on": {
                    "type": "string",
                    "format": "string",
                    "example": "2020-10-28T14:58:56.059Z"
                },
                "currency_code": {
                    "type": "string",
                    "format": "string",
                    "example": "RUB"
                },
                "date_expires": {
                    "type": "string",
                    "format": "string",
                    "example": "2020-10-28T14:58:56.059Z"
                },
                "description": {
                    "type": "string",
                    "format": "string",
                    "example": "Оригинальная сгущенка Рогачев"
                },
                "measure_unit_code": {
                    "type": "string",
                    "format": "string",
                    "example": "KG"
                },
                "offer_id": {
                    "type": "string",
                    "format": "string",
                    "example": "1d586b05-7b80-4a3a-bf2c-ce48169d4e85"
                },
                "offer_type": {
                    "type": "string",
                    "format": "string",
                    "example": "BUY"
                },
                "price": {
                    "type": "string",
                    "format": "int",
                    "example": "1000000"
                },
                "title": {
                    "type": "string",
                    "format": "string",
                    "example": "Сгущенка"
                },
                "user_id": {
                    "type": "string",
                    "format": "string",
                    "example": "1d586b05-7b80-4a3a-bf2c-ce48169d4e85"
                }
            }
        },
        "model.Record": {
            "type": "object",
            "properties": {
                "created_on": {
                    "type": "string",
                    "example": "2020-10-23 18:02:35.745565"
                },
                "id": {
                    "type": "string",
                    "example": "1d586b05-7b80-4a3a-bf2c-ce48169d4e85"
                }
            }
        },
        "model.Success": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        },
        "OAuth2AccessCode": {
            "type": "oauth2",
            "flow": "accessCode",
            "authorizationUrl": "https://example.com/oauth/authorize",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information"
            }
        },
        "OAuth2Application": {
            "type": "oauth2",
            "flow": "application",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
        },
        "OAuth2Implicit": {
            "type": "oauth2",
            "flow": "implicit",
            "authorizationUrl": "https://example.com/oauth/authorize",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
        },
        "OAuth2Password": {
            "type": "oauth2",
            "flow": "password",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "read": " Grants read access",
                "write": " Grants write access"
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
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample server celler server.",
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

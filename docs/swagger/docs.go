// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

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
        "/auth/login": {
            "post": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Авторизация",
                "parameters": [
                    {
                        "description": "Body params",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/github_com_Brotiger_per-painted_poker-backend_internal_module_auth_request.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный ответ.",
                        "schema": {
                            "$ref": "#/definitions/github_com_Brotiger_per-painted_poker-backend_internal_module_auth_response.Token"
                        }
                    },
                    "400": {
                        "description": "Не валидный запрос.",
                        "schema": {
                            "$ref": "#/definitions/github_com_Brotiger_per-painted_poker-backend_internal_shared_response.Error400"
                        }
                    },
                    "401": {
                        "description": "Не верное имя пользователя или пароль.",
                        "schema": {
                            "$ref": "#/definitions/github_com_Brotiger_per-painted_poker-backend_internal_shared_response.Error401"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера."
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Выход",
                "responses": {
                    "200": {
                        "description": "Успешный ответ."
                    },
                    "500": {
                        "description": "Ошибка сервера."
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Обновление токена",
                "parameters": [
                    {
                        "description": "Body params",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/github_com_Brotiger_per-painted_poker-backend_internal_module_auth_request.Refresh"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный ответ.",
                        "schema": {
                            "$ref": "#/definitions/github_com_Brotiger_per-painted_poker-backend_internal_module_auth_response.Token"
                        }
                    },
                    "400": {
                        "description": "Не валидный запрос.",
                        "schema": {
                            "$ref": "#/definitions/github_com_Brotiger_per-painted_poker-backend_internal_shared_response.Error400"
                        }
                    },
                    "401": {
                        "description": "Неверный или просроченный токен обновления.",
                        "schema": {
                            "$ref": "#/definitions/github_com_Brotiger_per-painted_poker-backend_internal_shared_response.Error401"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера."
                    }
                }
            }
        },
        "/game": {
            "get": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Получение списка игр",
                "parameters": [
                    {
                        "minimum": 0,
                        "type": "integer",
                        "example": 0,
                        "name": "from",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "test",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "example": 20,
                        "name": "size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный ответ.",
                        "schema": {
                            "$ref": "#/definitions/response.List"
                        }
                    },
                    "400": {
                        "description": "Не валидный запрос.",
                        "schema": {
                            "$ref": "#/definitions/github_com_Brotiger_per-painted_poker-backend_internal_shared_response.Error400"
                        }
                    },
                    "401": {
                        "description": "Невалидный токен.",
                        "schema": {
                            "$ref": "#/definitions/github_com_Brotiger_per-painted_poker-backend_internal_shared_response.Error401"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера."
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Создание игры",
                "parameters": [
                    {
                        "description": "Body params",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.Create"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный ответ.",
                        "schema": {
                            "$ref": "#/definitions/response.Create"
                        }
                    },
                    "400": {
                        "description": "Не валидный запрос.",
                        "schema": {
                            "$ref": "#/definitions/github_com_Brotiger_per-painted_poker-backend_internal_shared_response.Error400"
                        }
                    },
                    "401": {
                        "description": "Невалидный токен.",
                        "schema": {
                            "$ref": "#/definitions/github_com_Brotiger_per-painted_poker-backend_internal_shared_response.Error401"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера."
                    }
                }
            }
        },
        "/game/start": {
            "post": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Запуск игры",
                "responses": {
                    "200": {
                        "description": "Успешный ответ."
                    },
                    "401": {
                        "description": "Невалидный токен."
                    },
                    "500": {
                        "description": "Ошибка сервера."
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_Brotiger_per-painted_poker-backend_internal_module_auth_request.Login": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "username": {
                    "type": "string",
                    "example": "username"
                }
            }
        },
        "github_com_Brotiger_per-painted_poker-backend_internal_module_auth_request.Refresh": {
            "type": "object",
            "properties": {
                "refreshToken": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2V4YW1wbGUuYXV0aDAuY29tLyIsImF1ZCI6Imh0dHBzOi8vYXBpLmV4YW1wbGUuY29tL2NhbGFuZGFyL3YxLyIsInN1YiI6InVzcl8xMjMiLCJpYXQiOjE0NTg3ODU3OTYsImV4cCI6MTQ1ODg3MjE5Nn0.CA7eaHjIHz5NxeIJoFK9krqaeZrPLwmMmgI_XiQiIkQ"
                }
            }
        },
        "github_com_Brotiger_per-painted_poker-backend_internal_module_auth_response.Token": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2V4YW1wbGUuYXV0aDAuY29tLyIsImF1ZCI6Imh0dHBzOi8vYXBpLmV4YW1wbGUuY29tL2NhbGFuZGFyL3YxLyIsInN1YiI6InVzcl8xMjMiLCJpYXQiOjE0NTg3ODU3OTYsImV4cCI6MTQ1ODg3MjE5Nn0.CA7eaHjIHz5NxeIJoFK9krqaeZrPLwmMmgI_XiQiIkQ"
                },
                "refresh_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2V4YW1wbGUuYXV0aDAuY29tLyIsImF1ZCI6Imh0dHBzOi8vYXBpLmV4YW1wbGUuY29tL2NhbGFuZGFyL3YxLyIsInN1YiI6InVzcl8xMjMiLCJpYXQiOjE0NTg3ODU3OTYsImV4cCI6MTQ1ODg3MjE5Nn0.CA7eaHjIHz5NxeIJoFK9krqaeZrPLwmMmgI_XiQiIkQ"
                }
            }
        },
        "github_com_Brotiger_per-painted_poker-backend_internal_shared_response.Error400": {
            "type": "object",
            "properties": {
                "errors": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "github_com_Brotiger_per-painted_poker-backend_internal_shared_response.Error401": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "request.Create": {
            "type": "object",
            "properties": {
                "max_players": {
                    "type": "integer",
                    "maximum": 6,
                    "minimum": 3,
                    "example": 5
                },
                "name": {
                    "type": "string",
                    "example": "test"
                },
                "password": {
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "response.Create": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string",
                    "example": "507f1f77bcf86cd799439011"
                },
                "maxPlayers": {
                    "type": "integer",
                    "example": 5
                },
                "name": {
                    "type": "string",
                    "example": "test"
                },
                "ownerId": {
                    "type": "string",
                    "example": "507f1f77bcf86cd799439011"
                },
                "password": {
                    "type": "string",
                    "example": "123456"
                },
                "status": {
                    "type": "string",
                    "example": "created"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "507f1f77bcf86cd799439011"
                    ]
                }
            }
        },
        "response.Game": {
            "type": "object",
            "properties": {
                "countPlayers": {
                    "type": "integer",
                    "example": 3
                },
                "id": {
                    "type": "string",
                    "example": "507f1f77bcf86cd799439011"
                },
                "maxPlayers": {
                    "type": "integer",
                    "example": 4
                },
                "ownerId": {
                    "type": "string",
                    "example": "507f1f77bcf86cd799439011"
                },
                "status": {
                    "type": "string",
                    "example": "created"
                },
                "title": {
                    "type": "string",
                    "example": "test"
                },
                "withPassword": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "response.List": {
            "type": "object",
            "properties": {
                "games": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.Game"
                    }
                },
                "total": {
                    "type": "integer",
                    "example": 100
                }
            }
        }
    },
    "securityDefinitions": {
        "Authorization": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Core API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/app/v1/orders": {
            "post": {
                "description": "메뉴 주문을 할 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call Post Order, return posted id by json.",
                "parameters": [
                    {
                        "description": "RequestOrder JSON",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/protocol.RequestOrder"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/protocol.ApiResponse-any"
                        }
                    }
                }
            }
        },
        "/app/v1/stores": {
            "post": {
                "description": "가게정보를 등록 할 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call Post store, return posted id by json.",
                "parameters": [
                    {
                        "description": "RequestPostStore JSON",
                        "name": "store",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/protocol.RequestPostStore"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/protocol.ApiResponse-any"
                        }
                    }
                }
            }
        },
        "/app/v1/stores/menu": {
            "put": {
                "description": "메뉴를 수정할 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call Put menu, return updated count by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "menu-id",
                        "name": "menu-id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "RequestPostMenu JSON",
                        "name": "menu",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/protocol.RequestPostMenu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/protocol.ApiResponse-any"
                        }
                    }
                }
            },
            "post": {
                "description": "메뉴를 등록할 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call Post menu in store, return saved id by json.",
                "parameters": [
                    {
                        "description": "RequestPostMenu JSON",
                        "name": "menu",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/protocol.RequestPostMenu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/protocol.ApiResponse-any"
                        }
                    }
                }
            },
            "delete": {
                "description": "메뉴를 삭제할 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call Delete menu in store, return deleted count by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "store-id",
                        "name": "store-id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "menu-id",
                        "name": "menu-id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/protocol.ApiResponse-any"
                        }
                    }
                }
            }
        },
        "/app/v1/stores/swag/store": {
            "get": {
                "description": "특정 store 의 모든 정보를 스웨거 테스트를 위해 보여준다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call Get store, return store by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "store_id",
                        "name": "store_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/protocol.ApiResponse-entity_Store"
                        }
                    }
                }
            }
        },
        "/app/v1/users/join": {
            "post": {
                "description": "회원가입을 할 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call Post user, return saved id by json.",
                "parameters": [
                    {
                        "description": "RequestPostUser JSON",
                        "name": "menu",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/protocol.RequestPostUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/protocol.ApiResponse-any"
                        }
                    }
                }
            }
        },
        "/home/info": {
            "get": {
                "description": "App 에 대해 간략적인 정보를(소개) 제공해 준다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call App Information, return Info by json.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/protocol.ApiResponse-info_Info"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Address": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                },
                "zipCode": {
                    "type": "string"
                }
            }
        },
        "entity.BaseTime": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.Menu": {
            "type": "object",
            "properties": {
                "baseTime": {
                    "$ref": "#/definitions/entity.BaseTime"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "limitCount": {
                    "description": "OrderCount 총 주문수 --\u003e 주문내역으로 확인하자\nLimitCount 한정수량 ex) \"non\" , \"1\", \"10\"",
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "origin": {
                    "description": "Origin 원산지",
                    "type": "string"
                },
                "possible": {
                    "type": "boolean"
                },
                "price": {
                    "type": "integer"
                }
            }
        },
        "entity.Store": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/entity.Address"
                },
                "baseTime": {
                    "$ref": "#/definitions/entity.BaseTime"
                },
                "id": {
                    "type": "string"
                },
                "menu": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Menu"
                    }
                },
                "name": {
                    "type": "string"
                },
                "recommendMenus": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Menu"
                    }
                },
                "storePhone": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "info.Info": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "blog": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "spec": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "protocol.ApiResponse-any": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "protocol.ApiResponse-entity_Store": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/entity.Store"
                },
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "protocol.ApiResponse-info_Info": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/info.Info"
                },
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "protocol.RequestAddress": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                },
                "zip_code": {
                    "type": "string"
                }
            }
        },
        "protocol.RequestOrder": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "string"
                },
                "menu_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "ordered_addr": {
                    "$ref": "#/definitions/protocol.RequestAddress"
                },
                "price": {
                    "type": "integer"
                },
                "status": {
                    "description": "Status 접수중/접수취소/추가접수/접수중/조리중/배달중/배달완료",
                    "type": "string"
                },
                "store_id": {
                    "type": "string"
                }
            }
        },
        "protocol.RequestPostMenu": {
            "type": "object",
            "required": [
                "name",
                "origin",
                "possible",
                "price",
                "store_id"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "limit_count": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "origin": {
                    "type": "string"
                },
                "possible": {
                    "type": "boolean"
                },
                "price": {
                    "type": "integer"
                },
                "store_id": {
                    "type": "string"
                }
            }
        },
        "protocol.RequestPostStore": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/protocol.RequestAddress"
                },
                "name": {
                    "type": "string"
                },
                "store_phone": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "protocol.RequestPostUser": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "nic_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

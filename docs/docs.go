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
        "/app/v1/orders/order": {
            "get": {
                "description": "특정 주문기록을 볼 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "주문기록"
                ],
                "summary": "call Get order-record, return order-record by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "order-id",
                        "name": "order-id",
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
            },
            "post": {
                "description": "메뉴 주문을 할 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "주문기록"
                ],
                "summary": "call Post Order, return posted id by json.",
                "parameters": [
                    {
                        "description": "RequestOrder JSON",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RequestOrder"
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
        "/app/v1/orders/order/customer": {
            "put": {
                "description": "사용자가 주문을 변경 할 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "주문기록"
                ],
                "summary": "call Put order records in customer, return updated count by json.",
                "parameters": [
                    {
                        "description": "RequestPutCustomerOrder",
                        "name": "RequestPutCustomerOrder",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RequestPutCustomerOrder"
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
        "/app/v1/orders/order/price": {
            "get": {
                "description": "선택한 메뉴들의 총 가격을 알 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "주문기록"
                ],
                "summary": "call Get selected menus total price, return total price by json.",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "name": "menu_ids",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "store_id",
                        "in": "query"
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
        "/app/v1/orders/order/store": {
            "put": {
                "description": "가게에서 주문 상태를 변경 할 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "주문기록"
                ],
                "summary": "call Put order records in store, return updated count by json.",
                "parameters": [
                    {
                        "description": "RequestPutStoreOrder",
                        "name": "RequestPutStoreOrder",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RequestPutStoreOrder"
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
        "/app/v1/orders/pages/customer": {
            "get": {
                "description": "특정 사용자의 주문기록들을 볼 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "주문기록"
                ],
                "summary": "call Get sorted pages customer order records, return order records by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "customer-id",
                        "name": "customer-id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "content_count",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "current_page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort_name",
                        "in": "query"
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
        "/app/v1/orders/pages/store": {
            "get": {
                "description": "특정 가게의 주문기록들을 볼 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "주문기록"
                ],
                "summary": "call Get sorted pages store order records, return order records by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "store-id",
                        "name": "store-id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "content_count",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "current_page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort_name",
                        "in": "query"
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
        "/app/v1/reviews/customer": {
            "get": {
                "description": "특정 사용자의 리뷰들을 볼 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "메뉴리뷰"
                ],
                "summary": "call Get sorted page menu reviews, return sorted page menu reviews by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "customer-id",
                        "name": "customer-id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "content_count",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "current_page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort_name",
                        "in": "query"
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
        "/app/v1/reviews/menu": {
            "get": {
                "description": "특정 메뉴의 리뷰들을 볼 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "메뉴리뷰"
                ],
                "summary": "call Get sorted page menu reviews, return sorted page menu reviews by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "menu-id",
                        "name": "menu-id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "content_count",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "current_page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort_name",
                        "in": "query"
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
        "/app/v1/reviews/review": {
            "post": {
                "description": "메뉴 리뷰를 작성 할 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "메뉴리뷰"
                ],
                "summary": "call Post menu review, return saved id by json.",
                "parameters": [
                    {
                        "description": "RequestPostReview JSON",
                        "name": "RequestPostReview",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RequestPostReview"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/protocol.ApiResponse-any"
                        }
                    }
                }
            }
        },
        "/app/v1/stores": {
            "get": {
                "description": "가게들 정보를 보여준다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "가게"
                ],
                "summary": "call Get store pages, return store pages data by json.",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "content_count",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "current_page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort_name",
                        "in": "query"
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
        "/app/v1/stores/store": {
            "get": {
                "description": "특정 가게의 정보를 보여준다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "가게"
                ],
                "summary": "call Get store, return store by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "store_id",
                        "name": "store-id",
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
            },
            "put": {
                "description": "가게를 수정할 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "가게"
                ],
                "summary": "call Put store, return modify count by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "store-id",
                        "name": "store-id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "RequestPutStore JSON",
                        "name": "RequestPutStore",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RequestPutStore"
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
                "description": "가게정보를 등록 할 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "가게"
                ],
                "summary": "call Post store, return posted id by json.",
                "parameters": [
                    {
                        "description": "RequestPostStore JSON",
                        "name": "RequestPostStore",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RequestPostStore"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/protocol.ApiResponse-any"
                        }
                    }
                }
            }
        },
        "/app/v1/stores/store/menus": {
            "get": {
                "description": "특정 가게 메뉴 리스트를 보여준다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "가게"
                ],
                "summary": "call Get sorted menu page, return sorted menu pages data by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "store-id",
                        "name": "store-id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "content_count",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "current_page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort_name",
                        "in": "query"
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
        "/app/v1/stores/store/menus/menu": {
            "put": {
                "description": "메뉴를 수정할 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "가게"
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
                        "description": "RequestMenu JSON",
                        "name": "RequestMenu",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RequestMenu"
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
                "tags": [
                    "가게"
                ],
                "summary": "call Post menu, return saved id by json.",
                "parameters": [
                    {
                        "description": "RequestMenu JSON",
                        "name": "RequestMenu",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RequestMenu"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
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
                "tags": [
                    "가게"
                ],
                "summary": "call Delete menu, return deleted count by json.",
                "parameters": [
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
        "/app/v1/stores/store/recommends": {
            "get": {
                "description": "특정 가게의 추천 메뉴 상세 정보 보여준다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "가게"
                ],
                "summary": "call Get store and recommend menus, return store and recommend menus data by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "store-id",
                        "name": "store-id",
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
        "/app/v1/users/user": {
            "get": {
                "description": "사용자 정보를 보여준다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "사용자정보"
                ],
                "summary": "call Get user, return user by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user-id",
                        "name": "user-id",
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
            },
            "put": {
                "description": "사용자 정보를 수정 할 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "사용자정보"
                ],
                "summary": "call Put user, return updated count by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user-id",
                        "name": "user-id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "RequestUser JSON",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RequestUser"
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
                "description": "회원가입을 할 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "사용자정보"
                ],
                "summary": "call Post user, return saved id by json.",
                "parameters": [
                    {
                        "description": "RequestUser JSON",
                        "name": "menu",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RequestUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/protocol.ApiResponse-any"
                        }
                    }
                }
            },
            "delete": {
                "description": "사용자 정보를 삭제 할 수 있다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "사용자정보"
                ],
                "summary": "call Delete user, return delete count by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user-id",
                        "name": "user-id",
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
        "/home/info": {
            "get": {
                "description": "App 에 대해 간략적인 정보를(소개) 제공해 준다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "App"
                ],
                "summary": "call App Information, return Info by json.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/protocol.ApiResponse-any"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "page.Sort": {
            "type": "object",
            "properties": {
                "direction": {
                    "type": "integer"
                },
                "sort_name": {
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
        "request.RequestAddress": {
            "type": "object",
            "required": [
                "detail",
                "street",
                "zip_code"
            ],
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
        "request.RequestMenu": {
            "type": "object",
            "required": [
                "description",
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
        "request.RequestOrder": {
            "type": "object",
            "required": [
                "customer_id",
                "menu_ids",
                "ordered_addr",
                "store_id"
            ],
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
                    "$ref": "#/definitions/request.RequestAddress"
                },
                "store_id": {
                    "type": "string"
                }
            }
        },
        "request.RequestPostReview": {
            "type": "object",
            "required": [
                "content",
                "customer_id",
                "menu_id",
                "rating",
                "store_id"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "string"
                },
                "menu_id": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                },
                "store_id": {
                    "type": "string"
                }
            }
        },
        "request.RequestPostStore": {
            "type": "object",
            "required": [
                "address",
                "name",
                "store_phone",
                "user_id"
            ],
            "properties": {
                "address": {
                    "$ref": "#/definitions/request.RequestAddress"
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
        "request.RequestPutCustomerOrder": {
            "type": "object",
            "required": [
                "customer_id",
                "menu_ids",
                "order_id",
                "ordered_addr",
                "store_id"
            ],
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
                "order_id": {
                    "type": "string"
                },
                "ordered_addr": {
                    "$ref": "#/definitions/request.RequestAddress"
                },
                "store_id": {
                    "type": "string"
                }
            }
        },
        "request.RequestPutStore": {
            "type": "object",
            "required": [
                "address",
                "name",
                "store_phone",
                "user_id"
            ],
            "properties": {
                "address": {
                    "$ref": "#/definitions/request.RequestAddress"
                },
                "name": {
                    "type": "string"
                },
                "recommend_menus": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "store_phone": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "request.RequestPutStoreOrder": {
            "type": "object",
            "required": [
                "order_id",
                "status"
            ],
            "properties": {
                "order_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "request.RequestUser": {
            "type": "object",
            "required": [
                "name",
                "nic_name",
                "password",
                "phone_number",
                "role"
            ],
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
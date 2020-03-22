// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-03-20 20:56:42.422538 -0700 PDT m=+0.078380640

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
        "/v1/market/charts": {
            "get": {
                "description": "Get the charts data from an market and coin/token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Market"
                ],
                "summary": "Get charts data for a specific coin",
                "operationId": "get_charts_data",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 60,
                        "description": "Coin ID",
                        "name": "coin",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Token ID",
                        "name": "token",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 1574483028,
                        "description": "Start timestamp",
                        "name": "time_start",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 64,
                        "description": "Max number of items in result prices array",
                        "name": "max_items",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "USD",
                        "description": "The currency to show charts",
                        "name": "currency",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/watchmarket.ChartData"
                        }
                    }
                }
            }
        },
        "/v1/market/info": {
            "get": {
                "description": "Get the charts coin info data from an market and coin/contract",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Market"
                ],
                "summary": "Get charts coin info data for a specific coin",
                "operationId": "get_charts_coin_info",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 60,
                        "description": "Coin ID",
                        "name": "coin",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Token ID",
                        "name": "token",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "USD",
                        "description": "The currency to show coin info in",
                        "name": "currency",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/watchmarket.ChartCoinInfo"
                        }
                    }
                }
            }
        },
        "/v1/market/ticker": {
            "post": {
                "description": "Get the ticker values from many market and coin/token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Market"
                ],
                "summary": "Get ticker values for a specific market",
                "operationId": "get_tickers",
                "parameters": [
                    {
                        "description": "Ticker",
                        "name": "tickers",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.TickerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/watchmarket.Tickers"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Coin": {
            "type": "object",
            "properties": {
                "coin": {
                    "type": "integer"
                },
                "token_id": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "api.TickerRequest": {
            "type": "object",
            "properties": {
                "assets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Coin"
                    }
                },
                "currency": {
                    "type": "string"
                }
            }
        },
        "watchmarket.ChartCoinInfo": {
            "type": "object",
            "properties": {
                "circulating_supply": {
                    "type": "number"
                },
                "info": {
                    "type": "object",
                    "$ref": "#/definitions/watchmarket.CoinInfo"
                },
                "market_cap": {
                    "type": "number"
                },
                "total_supply": {
                    "type": "number"
                },
                "volume_24": {
                    "type": "number"
                }
            }
        },
        "watchmarket.ChartData": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "prices": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/watchmarket.ChartPrice"
                    }
                }
            }
        },
        "watchmarket.ChartPrice": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "watchmarket.CoinInfo": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "explorer": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "short_description": {
                    "type": "string"
                },
                "socials": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/watchmarket.SocialLink"
                    }
                },
                "source_code": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                },
                "white_paper": {
                    "type": "string"
                }
            }
        },
        "watchmarket.SocialLink": {
            "type": "object",
            "properties": {
                "handle": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "watchmarket.Ticker": {
            "type": "object",
            "properties": {
                "coin": {
                    "type": "integer"
                },
                "coin_name": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                },
                "last_update": {
                    "type": "string"
                },
                "price": {
                    "type": "object",
                    "$ref": "#/definitions/watchmarket.TickerPrice"
                },
                "token_id": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "watchmarket.TickerPrice": {
            "type": "object",
            "properties": {
                "change_24h": {
                    "type": "number"
                },
                "currency": {
                    "type": "string"
                },
                "provider": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "watchmarket.Tickers": {
            "type": "array",
            "items": {
                "$ref": "#/definitions/watchmarket.Ticker"
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

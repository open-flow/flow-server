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
        "/batch": {
            "post": {
                "parameters": [
                    {
                        "description": "save request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/batch.SaveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/batch.SaveResponse"
                        }
                    }
                }
            },
            "delete": {
                "parameters": [
                    {
                        "description": "delete request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/batch.DeleteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/callback": {
            "post": {
                "parameters": [
                    {
                        "description": "callback",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/call.CallbackRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/call.CallbackResponse"
                        }
                    }
                }
            }
        },
        "/connection": {
            "post": {
                "parameters": [
                    {
                        "description": "connection",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/graph.DBConnection"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/graph.DBConnection"
                        }
                    }
                }
            },
            "delete": {
                "parameters": [
                    {
                        "description": "connection id",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/graph.GraphObject"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/endpoint": {
            "post": {
                "parameters": [
                    {
                        "description": "endpoint",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/endpoint.DBEndpoint"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoint.DBEndpoint"
                        }
                    }
                }
            },
            "delete": {
                "parameters": [
                    {
                        "description": "endpoint id",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.ProjectModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/endpoint/list": {
            "get": {
                "parameters": [
                    {
                        "type": "integer",
                        "description": "project id",
                        "name": "projectId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoint.Container"
                        }
                    }
                }
            }
        },
        "/event-card": {
            "post": {
                "parameters": [
                    {
                        "description": "connection",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/graph.DBEventCard"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/graph.DBEventCard"
                        }
                    }
                }
            },
            "delete": {
                "parameters": [
                    {
                        "description": "event card id",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/graph.GraphObject"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/find-active": {
            "post": {
                "parameters": [
                    {
                        "description": "active event",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/search.FindActiveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/search.FindActiveResponse"
                        }
                    }
                }
            }
        },
        "/graph": {
            "get": {
                "parameters": [
                    {
                        "type": "integer",
                        "description": "project id",
                        "name": "projectId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "graph id",
                        "name": "graphId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/graph.DBGraph"
                        }
                    }
                }
            },
            "post": {
                "parameters": [
                    {
                        "description": "graph",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/graph.DBGraph"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/graph.DBGraph"
                        }
                    }
                }
            },
            "delete": {
                "parameters": [
                    {
                        "description": "graph id",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.ProjectModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/graph/list": {
            "get": {
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "description": "project ids",
                        "name": "projectId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/storage.ListGraphResponse"
                        }
                    }
                }
            }
        },
        "/node": {
            "post": {
                "parameters": [
                    {
                        "description": "node",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/graph.DBNode"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/graph.DBNode"
                        }
                    }
                }
            },
            "delete": {
                "parameters": [
                    {
                        "description": "node id",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/graph.GraphObject"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "batch.DeleteRequest": {
            "type": "object",
            "properties": {
                "cards": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "connections": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "nodes": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "projectId": {
                    "type": "integer"
                }
            }
        },
        "batch.SaveRequest": {
            "type": "object",
            "properties": {
                "cards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/graph.DBEventCard"
                    }
                },
                "connections": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/graph.DBConnection"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "nodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/graph.DBNode"
                    }
                },
                "projectId": {
                    "type": "integer"
                }
            }
        },
        "batch.SaveResponse": {
            "type": "object",
            "properties": {
                "cards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/graph.DBEventCard"
                    }
                },
                "connections": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/graph.DBConnection"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "nodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/graph.DBNode"
                    }
                },
                "projectId": {
                    "type": "integer"
                }
            }
        },
        "call.CallbackRequest": {
            "type": "object",
            "properties": {
                "context": {
                    "type": "object",
                    "additionalProperties": true
                },
                "event": {
                    "$ref": "#/definitions/graph.DataEvent"
                },
                "raw": {
                    "type": "object"
                }
            }
        },
        "call.CallbackResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "noExecutions": {
                    "type": "boolean"
                },
                "response": {
                    "type": "object"
                },
                "scheduled": {
                    "type": "boolean"
                },
                "timeout": {
                    "type": "boolean"
                }
            }
        },
        "common.ProjectModel": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "projectId": {
                    "type": "integer"
                }
            }
        },
        "endpoint.Container": {
            "type": "object",
            "properties": {
                "map": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/endpoint.Endpoint"
                    }
                }
            }
        },
        "endpoint.DBEndpoint": {
            "type": "object",
            "properties": {
                "headers": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "module": {
                    "type": "string"
                },
                "projectId": {
                    "type": "integer"
                },
                "query": {
                    "type": "string"
                },
                "uri": {
                    "type": "string"
                }
            }
        },
        "endpoint.Endpoint": {
            "type": "object",
            "properties": {
                "headers": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "module": {
                    "type": "string"
                },
                "uri": {
                    "type": "string"
                },
                "values": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "graph.DBConnection": {
            "type": "object",
            "properties": {
                "graphId": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "localId": {
                    "description": "Placement",
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "projectId": {
                    "type": "integer"
                },
                "sourceId": {
                    "type": "integer"
                },
                "sourcePort": {
                    "type": "string"
                },
                "targetId": {
                    "type": "integer"
                },
                "targetPort": {
                    "type": "string"
                },
                "ui": {
                    "type": "string"
                }
            }
        },
        "graph.DBEventCard": {
            "type": "object",
            "properties": {
                "contextId": {
                    "type": "string"
                },
                "contextType": {
                    "type": "string"
                },
                "graphId": {
                    "type": "integer"
                },
                "httpVote": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "initiatorId": {
                    "type": "string"
                },
                "initiatorType": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "ownerId": {
                    "type": "string"
                },
                "ownerType": {
                    "type": "string"
                },
                "platform": {
                    "type": "string"
                },
                "projectId": {
                    "type": "integer"
                },
                "resourceId": {
                    "type": "string"
                },
                "resourceType": {
                    "type": "string"
                },
                "slidePort": {
                    "type": "string"
                },
                "staticId": {
                    "type": "string"
                },
                "staticType": {
                    "type": "string"
                },
                "targetId": {
                    "description": "Placement",
                    "type": "integer"
                },
                "ui": {
                    "type": "string"
                }
            }
        },
        "graph.DBGraph": {
            "type": "object",
            "properties": {
                "cards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/graph.DBEventCard"
                    }
                },
                "connections": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/graph.DBConnection"
                    }
                },
                "counter": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "nodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/graph.DBNode"
                    }
                },
                "projectId": {
                    "type": "integer"
                },
                "ui": {
                    "type": "string"
                }
            }
        },
        "graph.DBNode": {
            "type": "object",
            "properties": {
                "arguments": {
                    "type": "string"
                },
                "function": {
                    "type": "string"
                },
                "graphId": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "localId": {
                    "description": "Placement",
                    "type": "integer"
                },
                "module": {
                    "description": "Invocation",
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "projectId": {
                    "type": "integer"
                },
                "ui": {
                    "type": "string"
                }
            }
        },
        "graph.DataEvent": {
            "type": "object",
            "properties": {
                "contextId": {
                    "type": "string"
                },
                "contextType": {
                    "type": "string"
                },
                "initiatorId": {
                    "type": "string"
                },
                "initiatorType": {
                    "type": "string"
                },
                "ownerId": {
                    "type": "string"
                },
                "ownerType": {
                    "type": "string"
                },
                "platform": {
                    "type": "string"
                },
                "resourceId": {
                    "type": "string"
                },
                "resourceType": {
                    "type": "string"
                },
                "staticId": {
                    "type": "string"
                },
                "staticType": {
                    "type": "string"
                }
            }
        },
        "graph.GraphObject": {
            "type": "object",
            "properties": {
                "graphId": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "projectId": {
                    "type": "integer"
                }
            }
        },
        "search.ActiveGraph": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/graph.DBEventCard"
                    }
                },
                "graph": {
                    "$ref": "#/definitions/graph.DBGraph"
                }
            }
        },
        "search.FindActiveRequest": {
            "type": "object",
            "properties": {
                "contextId": {
                    "type": "string"
                },
                "contextType": {
                    "type": "string"
                },
                "initiatorId": {
                    "type": "string"
                },
                "initiatorType": {
                    "type": "string"
                },
                "ownerId": {
                    "type": "string"
                },
                "ownerType": {
                    "type": "string"
                },
                "platform": {
                    "type": "string"
                },
                "resourceId": {
                    "type": "string"
                },
                "resourceType": {
                    "type": "string"
                },
                "staticId": {
                    "type": "string"
                },
                "staticType": {
                    "type": "string"
                }
            }
        },
        "search.FindActiveResponse": {
            "type": "object",
            "properties": {
                "graphs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/search.ActiveGraph"
                    }
                }
            }
        },
        "storage.ListGraphResponse": {
            "type": "object",
            "properties": {
                "graphs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/graph.DBGraph"
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

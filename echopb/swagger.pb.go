package echopb 

const (
swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "service.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/echo": {
      "post": {
        "operationId": "Echo",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/echopbEchoMessage"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/echopbEchoMessage"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/stream-echo": {
      "post": {
        "operationId": "StreamEcho",
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/echopbStreamEchoMessage"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "description": "(streaming inputs)",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/echopbStreamEchoMessage"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    }
  },
  "definitions": {
    "echopbEchoMessage": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string"
        }
      }
    },
    "echopbStreamEchoMessage": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string"
        }
      }
    }
  }
}
`
)

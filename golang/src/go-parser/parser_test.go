package main

import "encoding/json"

type RespBody struct {
	RequestID string      `json:"request_id"`
	Message   string      `json:"message"`
	Code      int         `json:"code"`
	ErrorCode string      `json:"error_code,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Data2     json.RawMessage
}

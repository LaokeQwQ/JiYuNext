package core

import (
	"encoding/json"
	"fmt"
)

const (
	CodeOK           = 0
	CodeInvalidInput = 1001
)

type Request struct {
	Action  string          `json:"action"`
	Payload json.RawMessage `json:"payload,omitempty"`
	TraceID string          `json:"trace_id,omitempty"`
}

type Response struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result,omitempty"`
	TraceID string          `json:"trace_id,omitempty"`
}

func Init(configJSON string) Response {
	if configJSON == "" {
		return ok(`{"initialized":true}`)
	}
	if !json.Valid([]byte(configJSON)) {
		return fail(CodeInvalidInput, "invalid config_json")
	}
	return ok(`{"initialized":true}`)
}

func Execute(actionJSON string) Response {
	var req Request
	if err := json.Unmarshal([]byte(actionJSON), &req); err != nil {
		return fail(CodeInvalidInput, "invalid action_json")
	}
	if req.Action == "" {
		return fail(CodeInvalidInput, "action is required")
	}
	if req.TraceID == "" {
		req.TraceID = "trace-missing"
	}
	result := fmt.Sprintf(`{"accepted":true,"action":"%s"}`, req.Action)
	resp := ok(result)
	resp.TraceID = req.TraceID
	return resp
}

func Shutdown() Response {
	return ok(`{"shutdown":true}`)
}

func ok(result string) Response {
	return Response{
		Code:    CodeOK,
		Message: "ok",
		Result:  json.RawMessage(result),
	}
}

func fail(code int, message string) Response {
	return Response{
		Code:    code,
		Message: message,
	}
}

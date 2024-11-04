package jsonrpc

import (
	"encoding/json"
)

// Incoming JSON-RPC request sent to Joute's endpoint
type Request struct {
	Id     string `json:"id"`
	Method string `json:"method"`
}

func UnmarshallBytes(payload []byte) (*Request, error) {
	var req Request
	err := json.Unmarshal(payload, &req)
	return &req, err
}

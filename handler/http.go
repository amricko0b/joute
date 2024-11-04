package handler

import (
	"bytes"
	"io"
	"net/http"

	"github.com/amricko0b/joute/jsonrpc"
)

// JouteHandler acts as a requests proxy
type JouteHandler struct {
	TargetScheme   string
	TargetHostPort string

	Client *http.Client
}

// ServeHTTP contains proxy logic.
// Request is modified with utility headers.
func (h *JouteHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	rpcReq, err := unmarshallBodyIntact(req)
	if err != nil {
		http.Error(rw, "Not a JSON-RPC request", http.StatusBadRequest)
		return
	}

	h.modifyRequest(req, rpcReq)
	resp, err := h.Client.Do(req)
	if err != nil {
		http.Error(rw, "Server Error", http.StatusInternalServerError)
		return
	}

	for key, values := range resp.Header {
		for _, val := range values {
			rw.Header().Add(key, val)
		}
	}

	defer resp.Body.Close()
	rw.WriteHeader(resp.StatusCode)
	io.Copy(rw, resp.Body)
}

// modifyRequest prepares original request before being proxied to target.
// It sets JSON-RPC headers to original request and changes desired host to target.
//
// Some of request's properties are not expected to be used from in client context so we erase them.
func (h *JouteHandler) modifyRequest(req *http.Request, rpcReq *jsonrpc.Request) {
	req.Header.Set("X-Rpc-Method", rpcReq.Method)
	req.Header.Set("X-Rpc-Id", rpcReq.Id)

	req.URL.Scheme = h.TargetScheme
	req.URL.Host = h.TargetHostPort

	req.Host = ""
	req.RequestURI = ""
}

// unmarshallBodyIntact does common JSON-RPC request unmarshalling.
// In order to investigate JSON-RPC request body must be read.
// But reading the body erases it from request, so we must buffer it as byte slice and then write back to request.
func unmarshallBodyIntact(req *http.Request) (*jsonrpc.Request, error) {

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	rpcReq, err := jsonrpc.UnmarshallBytes(body)
	if err != nil {
		return nil, err
	}

	req.Body = io.NopCloser(bytes.NewBuffer(body))
	return rpcReq, nil
}

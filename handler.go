package joute

import (
	"bytes"
	"io"
	"net/http"

	"github.com/amricko0b/joute/jsonrpc"
)

// Handler provides main Joute's HTTP handling logic hence must be exported by HTTP server.
// Captures Client's request and proxies it to downstream.
type Handler struct {
	RequestModifiers  []RequestModifier
	ResponseModifiers []ResponseModifier

	Client *http.Client
}

func (h *Handler) ServeHTTP(clientResponse http.ResponseWriter, clientRequest *http.Request) {

	clientMessage, err := unmarshallBodyIntact(clientRequest)
	if err != nil {
		// TODO specification is violated here...
		// TODO when bad JSON-RPC received from client we must answer with appropriate JSON-RPC response ("Parse error" or "Invalid Request")
		http.Error(clientResponse, "Joute: not a JSON-RPC request", http.StatusBadRequest)
		return
	}

	for _, modifier := range h.RequestModifiers {
		modifier.Modify(clientMessage, clientRequest)
	}

	downstreamResponse, err := h.Client.Do(clientRequest)
	if err != nil {
		http.Error(clientResponse, "Joute: downstream interaction error", http.StatusInternalServerError)
		return
	}

	defer downstreamResponse.Body.Close()
	for _, modifier := range h.ResponseModifiers {
		modifier.Modify(downstreamResponse, clientResponse)
	}
}

// After reading the body it won't be available anymore (actually it is read directly from socket) - so we must buffer it and populate back.
// Apart from this this function does regular JSON-RPC unmarshalling.
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

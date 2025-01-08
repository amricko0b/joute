package joute

import (
	"net/http"

	"github.com/amricko0b/joute/jsonrpc"
)

// RequestModifier is used BEFORE proxying and modifies Client's request to the Downstream.
// In order to execute modification JSON-RPC Request message must be parsed from Client request body.
// An implementation may propagate values from clientMessage to downstreamRequest as well as some arbitrary values.
// Due to mutable nature of Modifier the order of modifications must be noted.
type RequestModifier interface {
	Modify(clientMessage *jsonrpc.Request, downstreamRequest *http.Request) error
}

// ResponseModifier is used AFTER proxying and modifies Downstream's response to Client.
// An implementation may propagate values from downstreamResponse to clientResponse as well as some arbitrary values.
// Due to mutable nature of Modifier the order of modifications must be noted.
type ResponseModifier interface {
	Modify(downstreamResponse *http.Response, clientResponse http.ResponseWriter) error
}

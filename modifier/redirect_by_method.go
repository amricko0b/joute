package modifier

import (
	"fmt"
	"net/http"

	"github.com/amricko0b/joute/jsonrpc"
)

// RedirectByMethod redirects Client's request to Downstream (changes request address).
// It uses JSON-RPC method to construct new request path (method will be put as last fragment of path)
type RedirectByMethod struct {
	DownstreamScheme   string
	DownstreamHostPort string
}

func (r *RedirectByMethod) Modify(clientMessage *jsonrpc.Request, downstreamRequest *http.Request) error {

	downstreamRequest.URL.Path = fmt.Sprintf("%s/%s", downstreamRequest.URL.Path, clientMessage.Method)

	downstreamRequest.URL.Scheme = r.DownstreamScheme
	downstreamRequest.URL.Host = r.DownstreamHostPort

	// This is mandatory when proxying HTTP request to another server!
	downstreamRequest.Host = ""
	downstreamRequest.RequestURI = ""

	return nil
}

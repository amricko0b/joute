package modifier

import (
	"net/http"

	"github.com/amricko0b/joute/jsonrpc"
)

// Redirect Client's request to Downstream
type Redirect struct {
	DownstreamScheme   string
	DownstreamHostPort string
}

func (r *Redirect) Modify(clientMessage *jsonrpc.Request, downstreamRequest *http.Request) error {

	downstreamRequest.URL.Scheme = r.DownstreamScheme
	downstreamRequest.URL.Host = r.DownstreamHostPort

	// This is mandatory when proxying HTTP request to another server!
	downstreamRequest.Host = ""
	downstreamRequest.RequestURI = ""

	return nil
}

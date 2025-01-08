package modifier

import (
	"net/http"

	"github.com/amricko0b/joute/jsonrpc"
)

const (
	RpcMethodHeader = "X-Rpc-Method"
	RpcIdHeader     = "X-Rpc-Id"
)

// AddHeaders enriches Client's request with RPC headers before sending to Downstream
type AddHeaders struct{}

func (m *AddHeaders) Modify(clientMessage *jsonrpc.Request, downstreamRequest *http.Request) error {

	downstreamRequest.Header.Set(RpcMethodHeader, clientMessage.Method)
	downstreamRequest.Header.Set(RpcIdHeader, clientMessage.Id)

	return nil
}

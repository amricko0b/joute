package modifier_test

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/amricko0b/joute/jsonrpc"
	"github.com/amricko0b/joute/modifier"
)

func TestHeadersAreAdded(t *testing.T) {
	request := http.Request{Header: make(http.Header)}
	call := jsonrpc.Request{
		Id:     uuid.NewString(),
		Method: "do_something",
	}

	m := modifier.AddHeaders{}
	m.Modify(&call, &request)

	assert.NotEmpty(t, request.Header.Get(modifier.RpcMethodHeader))
	assert.NotEmpty(t, request.Header.Get(modifier.RpcIdHeader))
}

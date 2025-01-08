package modifier_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/amricko0b/joute/jsonrpc"
	"github.com/amricko0b/joute/modifier"
)

func TestRequestRedirected(t *testing.T) {

	request := http.Request{URL: &url.URL{}}
	call := jsonrpc.Request{}

	m := modifier.Redirect{DownstreamScheme: "https", DownstreamHostPort: "example.com:443"}
	m.Modify(&call, &request)

	assert.Equal(t, m.DownstreamScheme, request.URL.Scheme)
	assert.Equal(t, m.DownstreamHostPort, request.URL.Host)

	assert.Empty(t, request.Host)
	assert.Empty(t, request.RequestURI)
}

package modifier_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/amricko0b/joute/jsonrpc"
	"github.com/amricko0b/joute/modifier"
)

func TestMethodIsAddedAsPathPostfix(t *testing.T) {

	u := url.URL{
		Scheme: "https",
		Host:   "joute.svc:443",
		Path:   "/api/v1/jsonrpc",
	}

	request := http.Request{URL: &u}
	call := jsonrpc.Request{
		Method: "do_something",
	}

	m := modifier.RedirectByMethod{
		DownstreamScheme:   "https",
		DownstreamHostPort: "example.com:443",
	}

	m.Modify(&call, &request)

	assert.Equal(t, "/api/v1/jsonrpc/do_something", request.URL.Path)

	assert.Equal(t, m.DownstreamScheme, request.URL.Scheme)
	assert.Equal(t, m.DownstreamHostPort, request.URL.Host)

	assert.Empty(t, request.Host)
	assert.Empty(t, request.RequestURI)
}

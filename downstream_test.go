package joute_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/amricko0b/joute"
)

func TestDownstreamsLoad(t *testing.T) {
	app, err := joute.LoadAppWithConfigFrom(joute.ConfigFileLocation("./configs"))
	assert.NoError(t, err)
	assert.NotNil(t, app)

	assert.Contains(t, app.Downstreams, "primarchs")
	assert.Contains(t, app.Downstreams, "legions")

	legions := app.Downstreams["legions"]
	assert.Equal(t, legions.URL.Scheme, "https")
	assert.Equal(t, legions.URL.Host, "legions.svc")
	assert.Equal(t, joute.DownstreamTimeout(10*time.Second), legions.Timeout)

	primarchs := app.Downstreams["primarchs"]
	assert.Equal(t, primarchs.URL.Scheme, "https")
	assert.Equal(t, primarchs.URL.Host, "primarchs.svc:443")
	assert.Equal(t, joute.DownstreamTimeout(1*time.Minute), primarchs.Timeout)
}

func TestDownstreamMayBeCalled(t *testing.T) {

	srv := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprint(resp, "Pong!")
		resp.WriteHeader(200)
	}))

	defer srv.Close()

	addr, err := url.Parse(srv.URL)
	assert.NoError(t, err)

	ds := joute.Downstream{
		URL: (*joute.DownstreamURL)(addr), Timeout: joute.DownstreamTimeout(5 * time.Second),
	}

	resp, err := ds.Call(httptest.NewRequest(http.MethodPost, "http://pingpong.svc/", bytes.NewBufferString("Ping!")))
	assert.NoError(t, err)

	defer resp.Body.Close()
	payload, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, "Pong!", string(payload))
	assert.Equal(t, 200, resp.StatusCode)
}

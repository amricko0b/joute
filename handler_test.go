package joute_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/amricko0b/joute"
	"github.com/amricko0b/joute/modifier"
)

// Very clumsy test - must be rewritten someday...
func TestHandlerAppliesRedirection(t *testing.T) {

	downstream := startDownstream()
	defer downstream.Close()

	joute := startJoute(downstream)
	defer joute.Close()

	file, err := os.Open("./jsonrpc/primarch_create.json")
	if err != nil {
		t.Error(err)
	}

	response, err := http.Post(joute.URL, "application/json", file)
	assert.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, 200, response.StatusCode)

	var msg map[string]interface{}
	payload, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(payload, &msg)
	assert.NoError(t, err)

	assert.Equal(t, "2.0", msg["jsonrpc"])
	assert.Equal(t, "1130ac93-2d76-49e1-bc32-7787436d81be", msg["id"])
}

func startDownstream() *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			res := make(map[string]string)
			res["ok"] = "true"

			m := make(map[string]interface{})
			m["jsonrpc"] = "2.0"
			m["id"] = "1130ac93-2d76-49e1-bc32-7787436d81be"
			m["result"] = res

			json.NewEncoder(w).Encode(m)
		}),
	)
}

func startJoute(downstream *httptest.Server) *httptest.Server {
	return httptest.NewServer(&joute.Handler{
		RequestModifiers: []joute.RequestModifier{
			&modifier.Redirect{
				DownstreamScheme:   "http",
				DownstreamHostPort: strings.Replace(downstream.URL, "http://", "", 1),
			},
		},
		ResponseModifiers: []joute.ResponseModifier{&modifier.OriginalResponseBody{}},
		Client:            &http.Client{},
	})
}

package handler_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/assert"
	wiremock "github.com/wiremock/wiremock-testcontainers-go"

	"github.com/amricko0b/joute/handler"
)

func TestServeHTTP(t *testing.T) {

	ctx := context.Background()

	wm, err := wiremock.RunContainerAndStopOnCleanup(ctx,
		t,
		wiremock.WithMappingFile("primarch_create", "mappings/primarch_create.json"),
	)
	if err != nil {
		t.Error(err)
	}

	httpPort, err := nat.NewPort("tcp", "8080")
	assert.NoError(t, err)

	containerHttpPort, err := wm.MappedPort(ctx, httpPort)
	assert.NoError(t, err)

	srv := httptest.NewServer(&handler.JouteHandler{
		TargetScheme:   "http",
		TargetHostPort: fmt.Sprintf("localhost:%d", containerHttpPort.Int()),
		Client:         &http.Client{},
	})
	defer srv.Close()

	file, err := os.Open("../jsonrpc/primarch_create.json")
	if err != nil {
		t.Error(err)
	}

	response, err := http.Post(srv.URL, "application/json", file)
	assert.NoError(t, err)
	defer response.Body.Close()

	assert.Equal(t, 200, response.StatusCode)

	payload, _ := io.ReadAll(response.Body)
	assert.Equal(t, "{\"jsonrpc\":\"2.0\",\"id\":\"1130ac93-2d76-49e1-bc32-7787436d81be\",\"result\":{\"ok\":true}}", string(payload))
}

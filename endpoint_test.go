package joute_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/amricko0b/joute"
)

func TestEndpointsLoad(t *testing.T) {
	app, err := joute.LoadAppWithConfigFrom(joute.ConfigFileLocation("./configs"))
	assert.NoError(t, err)
	assert.NotNil(t, app)

	assert.Contains(t, app.Endpoints, "/api/v1/primarch/jsonrpc")

	endpoint := app.Endpoints["/api/v1/primarch/jsonrpc"]
	assert.Equal(t, endpoint.Config.RouteTo, "primarchs")
	assert.Equal(t, endpoint.Config.Routing, joute.RoutingDirect)
}

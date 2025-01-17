package joute_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/amricko0b/joute"
)

func TestConfigMapLoads(t *testing.T) {

	configMap, err := joute.LoadConfigMap()
	assert.NoError(t, err)

	assert.Contains(t, configMap, "downstreams")
	assert.Contains(t, configMap, "endpoints")
}

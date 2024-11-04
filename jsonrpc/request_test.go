package jsonrpc_test

import (
	"io"
	"os"
	"testing"

	"github.com/amricko0b/joute/jsonrpc"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshallBytes(t *testing.T) {

	file, err := os.Open("./primarch_create.json")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	defer file.Close()
	payload, err := io.ReadAll(file)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	req, err := jsonrpc.UnmarshallBytes(payload)
	assert.NoError(t, err)
	assert.Equal(t, "primarch.create", req.Method)
}

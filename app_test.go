package joute_test

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/amricko0b/joute"
)

func TestAppRuns(t *testing.T) {
	app, err := joute.LoadAppWithConfigFrom(joute.ConfigFileLocation("./configs"))
	assert.NoError(t, err)
	go app.Run()

	isHttpServed := func() bool {
		address := net.JoinHostPort("127.0.0.1", fmt.Sprintf("%d", app.Port))
		_, err := net.DialTimeout("tcp", address, 3*time.Second)
		return err == nil
	}

	assert.Eventually(t, isHttpServed, 5*time.Second, 1*time.Second)
}

func TestAppLoadsWithConfigFileLocation(t *testing.T) {
	app, err := joute.LoadAppWithConfigFrom(joute.ConfigFileLocation("./configs"))
	assert.NoError(t, err)
	assert.NotNil(t, app)

	assert.Equal(t, 9000, app.Port)
}

func TestAppDoesNotLoadFromWorkingDirectoryDueToNoConfigFile(t *testing.T) {
	app, err := joute.LoadApp()

	assert.Error(t, err)
	assert.ErrorContains(t, err, "no such file or directory")
	assert.Nil(t, app)
}
